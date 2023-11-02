package main

import "fmt"

// Example of the Array of the Structure
//
//	type Addres struct {
//		Village       string
//		Post          string
//		policeStation string
//		PinCode       int
//		District      string
//		State         string
//	}
type Student struct {
	Name          string
	Roll          int
	Age           float64
	Email         string
	Village       string
	Post          string
	policeStation string
	PinCode       int
	District      string
	State         string
	//Ad   Addres
}

type College struct {
	CollegeNmae       string
	CollegePlace      string
	CollegeCode       int
	CollegeDepartment string
	Stu               []Student
	Ad                []Addres
}

func (c College) getinfo() string {
	fmt.Println("College Name -", c.CollegeNmae)
	fmt.Println("College Place -", c.CollegePlace)
	fmt.Println("College Code -", c.CollegeCode)
	fmt.Println("College Department -", c.CollegeDepartment)

	for _, i := range c.Stu {
		fmt.Println("Name -", i.Name)
		fmt.Println("Roll -", i.Roll)
		fmt.Println("Age -", i.Age)
		fmt.Println("Email -", i.Email)
		fmt.Println("Village -", i.Ad.Village)
		fmt.Println("Post -", i.Ad.Post)
		fmt.Println("Police -", i.Ad.policeStation)
		fmt.Println("Pin Code -", i.Ad.PinCode)
		fmt.Println("District -", i.Ad.District)
		fmt.Println("State -", i.Ad.State)
	}

	return "---------------"
}
func main() {
	c := College{
		CollegeNmae:       "Sachchidanand Sinha College Auranagabad (Bihar)",
		CollegePlace:      "Auranagabad",
		CollegeCode:       152,
		CollegeDepartment: "Computer Application",
		Stu: Student{
			Name:  "",
			Roll:  0,
			Age:   0,
			Email: "",
			Ad: Addres{
				Village:       "",
				Post:          "",
				policeStation: "",
				PinCode:       0,
				District:      "",
				State:         "",
			},
		},
	}

}
