package api

// import "github.com/IsaErtunga/companies/internal/ports"

// type Adapter struct {
// 	db    ports.DbPort
// 	arith ports.ArithmeticPort
// }

// func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
// 	return &Adapter{
// 		db:    db,
// 		arith: arith,
// 	}
// }

// func (apia Adapter) GetAddition(a, b int32) (int32, error) {
// 	ans, err := apia.arith.Addition(a, b)
// 	if err != nil {
// 		return 0, err
// 	}
// 	err = apia.db.AddToHistory(ans, "addition")
// 	if err != nil {
// 		return 0, err
// 	}

// 	return ans, err
// }

// func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
// 	ans, err := apia.arith.Subtraction(a, b)
// 	if err != nil {
// 		return 0, err
// 	}

// 	err = apia.db.AddToHistory(ans, "subtraction")
// 	if err != nil {
// 		return 0, err
// 	}

// 	return ans, err
// }

// func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
// 	ans, err := apia.arith.Multiplication(a, b)
// 	if err != nil {
// 		return 0, err
// 	}

// 	err = apia.db.AddToHistory(ans, "multiplication")
// 	if err != nil {
// 		return 0, err
// 	}

// 	return ans, err
// }

// func (apia Adapter) GetDivision(a, b int32) (int32, error) {
// 	ans, err := apia.arith.Division(a, b)
// 	if err != nil {
// 		return 0, err
// 	}

// 	err = apia.db.AddToHistory(ans, "division")
// 	if err != nil {
// 		return 0, err
// 	}

// 	return ans, err
// }
