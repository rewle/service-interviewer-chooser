package calendar

type Calendar struct{}

type Time struct{}

type Restrictions struct{}

type Interviewer struct{}

func (i *Interviewer) Avaliable(t Time) bool {
	return true
}

func (i *Interviewer) SetPriority(r Restrictions) bool {
	return true
}

func getInterviewverList() []*Interviewer {
	return []*Interviewer{}
}

func choose(t Time) []*Interviewer {
	interviewers := make([]*Interviewer, 0)
	for _, i := range getInterviewverList() {
		if i.Avaliable(t) {
			i.SetPriority(Restrictions{}) // Restrictions from ctx
			interviewers = append(interviewers, i)
		}
	}
	return interviewers
	// get list of interwievers
	// filter avaliability by Time
	// get priority by restrictions
	// return
}
