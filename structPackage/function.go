package structpackage

type Student struct {
	fname string
	lname string
	age   int
}

func (s *Student) Getfname () string {
	return s.fname
}

func (s *Student) Getlname () string {
	return s.lname
}

func (s *Student) Getage () int {
	return s.age
}

func (s *Student) Setfname (fname string)  {
	s.fname = fname
}

func (s *Student) Setlname (lname string)  {
	s.lname = lname
}

func (s *Student) Setage (age int) {
	s.age = age
}

