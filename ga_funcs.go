package goga

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
	s.fitness = defaultFitness
	s.config.maxNumOfSteps = defaultMaxNumOfSteps
	s.config.targetCost = defaultTargetCost
	s.config.stepsInterval = defaultStepsInterval
	s.config.selection.top = defaultSelectionTop
	s.config.selection.mutation = defaultSelectionMutation
	s.config.selection.random = float64(1) - defaultSelectionTop - defaultSelectionMutation
	s.config.reverse = defaultReverse
}

func (s *ga) init() error {
	return nil
}
