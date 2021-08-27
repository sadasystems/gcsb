package data

type (
	ThreadDataGenerator struct {
		prefixLength     int
		stringLength     int
		rowCount         int
		threadCount      int
		threadGenerators []*CombinedGenerator
	}

	ThreadDataGeneratorConfig struct {
		PrefixLength int
		StringLength int
		RowCount     int
		ThreadCount  int
	}
)

func NewThreadDataGenerator(cfg ThreadDataGeneratorConfig) (*ThreadDataGenerator, error) {
	ret := &ThreadDataGenerator{
		prefixLength: cfg.PrefixLength,
		stringLength: cfg.StringLength,
		threadCount:  cfg.ThreadCount,
		rowCount:     cfg.RowCount,
	}

	return ret, nil
}

func (s *ThreadDataGenerator) GetThreadGenerators() {
	rowsPerThread := s.rowCount / s.threadCount
	i := 0
	threads := make([]*CombinedGenerator, s.threadCount)
	for i < s.threadCount {
		gen, _ := NewCombinedGenerator(CombinedGeneratorConfig{
			Min:          i * rowsPerThread,
			Max:          (i+1)*rowsPerThread - 1,
			PrefixLength: s.prefixLength,
			StringLength: s.stringLength,
		})
		threads[i] = gen
		i++
	}
	s.threadGenerators = threads
}
