package dollar

type Dollar struct {
  amount int
}

func (dollar Dollar) Times(by int) int {
   return dollar.amount  *  by
}



