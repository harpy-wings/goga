package goga

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

func New(ops ...Option) (GA, error) {
	g := new(ga)

	g.loadDefaults()

	for _, fn := range ops {
		err := fn(g)
		if err != nil {
			return nil, err
		}
	}

	err := g.init()
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (s *ga) loadDefaults() {
	s.config.initialPopulation = defaultInitialPopulation
	s.population = defaultPopulation
	s.weightFunc = defaultWeightFunc
	s.config.maxNumOfSteps = defaultMaxNumOfSteps
	s.config.targetCost = defaultTargetCost
	s.config.stepsInterval = defaultStepsInterval
	s.config.selection.top = defaultSelectionTop
	s.config.selection.mutation = defaultSelectionMutation
	s.config.selection.random = float64(1) - defaultSelectionTop - defaultSelectionMutation
	s.config.numberOfThreads = defaultNumberOfThreads
}

func (s *ga) RuntimeBestResult() chan RunTimeResult {
	return s.runtimeResult
}

func (s *ga) init() error {
	s.runtimeError = make(chan error, 10)
	s.result = make(chan Model, 1)
	s.runtimeResult = make(chan RunTimeResult, 1000)
	return nil
}

// Start the Process
func (s *ga) Start() error {
	// generate the first population
	// calculate cost for each model ||Parallel
	// sort the generation
	// check stopping condition
	// if end respond on Queue
	// if not end, wait the interval then start over.
	go func() {
		err := s.generateRemindingGeneration(s.config.initialPopulation)
		if err != nil {
			s.runtimeError <- err
		}
		s.calculateCosts()
		s.sort()
		s.step = 0
		for {
			genLen := len(s.curetGeneration)
			nextPopulation := s.population(genLen, s.step, s.curetGeneration[0].cost, s.curetGeneration[genLen-1].cost)
			top := s.getTop(int(nextPopulation))
			mut := s.getMutation(int(nextPopulation))
			s.curetGeneration = append(top, mut...)
			err := s.generateRemindingGeneration(nextPopulation)
			if err != nil {
				s.runtimeError <- err
			}
			s.calculateCosts()
			s.sort()
			s.runtimeResult <- RunTimeResult{
				Model: s.curetGeneration[0].model,
				Cost:  s.curetGeneration[0].cost,
				Step:  s.step,
			}
			if s.curetGeneration[0].cost <= s.config.targetCost {
				break
			}
			if s.config.maxNumOfSteps != 0 && s.step > s.config.maxNumOfSteps {
				break
			}
			time.Sleep(s.config.stepsInterval)
			s.step++
		}
		s.result <- s.curetGeneration[0].model
	}()

	return nil
}

// todo add comment
func (s *ga) Result() (Model, error) {
	select {
	case res, ok := <-s.result:
		if !ok {
			return nil, ErrExecutionFailed("cannot reading data from a closed channel")
		}
		return res, nil

	case err := <-s.runtimeError:
		return nil, err
	}
}

func (s *ga) generateRemindingGeneration(nextPopulation int) error {
	for i := len(s.curetGeneration); i < nextPopulation; i++ {
		s.curetGeneration = append(s.curetGeneration, modelRecord{s.generator(), 0})
	}
	return nil
}

func (s *ga) calculateCosts() {
	C := make(chan int, len(s.curetGeneration))
	for i := range s.curetGeneration {
		C <- i
	}
	wg := &sync.WaitGroup{}

	wg.Add(s.config.numberOfThreads)

	for i := 0; i < s.config.numberOfThreads; i++ {
		go func() {
		InnerLoop:
			for {
				select {
				case j, ok := <-C:
					if !ok {
						break InnerLoop
					}
					cost := s.curetGeneration[j].model.Cost()
					s.Lock()
					s.curetGeneration[j].cost = cost
					s.Unlock()
				default:
					break InnerLoop
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(C)
}

func (s *ga) sort() {
	sort.Sort(s.curetGeneration)
}

func (s *ga) getTop(population int) []modelRecord {
	return s.curetGeneration[:int(s.config.selection.top*float64(population))]
}
func (s *ga) getMutation(population int) []modelRecord {
	//todo add check that the generation size is not grow too fast.
	var selectingItems []int
	List := s.curetGeneration[:int((s.config.selection.mutation+s.config.selection.top)*float64(population))]
	//fill up the weighted indexes
	for i, v := range List {
		w := s.weightFunc(i, v.cost)
		for j := 0; j < w; j++ {
			selectingItems = append(selectingItems, i)
		}
	}

	n := len(selectingItems) - 1
	targetLen := int(s.config.selection.mutation * float64(population))
	var res []modelRecord
	for i := 0; i < targetLen/2; i++ {
		A := List[selectingItems[rand.Intn(n)]].model
		B := List[selectingItems[rand.Intn(n)]].model
		res = append(res, modelRecord{A.Mutation(B), 0}, modelRecord{B.Mutation(A), 0})
	}
	return res
}
