package main

import "fmt"

// Student form
type Addres struct {
	Village       string
	Post          string
	policeStation string
	PinCode       int
	District      string
	State         string
}
type CollegeDetails struct {
	CollegeName       string
	CollegeCode       int
	CollegeDepartment string
}
type StDetails struct {
	Name     string
	Roll     int
	Age      int
	Ad       Addres
	Colleges CollegeDetails
}
type FormDetails struct {
	Students []StDetails
}

func (f FormDetails) studentForm() string {
	for _, i := range f.Students {
		fmt.Println()
		fmt.Println("Name -", i.Name)
		fmt.Println("Roll -", i.Roll)
		fmt.Println("Age -", i.Age)
		fmt.Println("Village -", i.Ad.Village)
		fmt.Println("Post -", i.Ad.Post)
		fmt.Println("Police -", i.Ad.policeStation)
		fmt.Println("Pin Code -", i.Ad.PinCode)
		fmt.Println("District -", i.Ad.District)
		fmt.Println("State -", i.Ad.State)
		fmt.Println("College Name -", i.Colleges.CollegeName)
		fmt.Println("College Code -", i.Colleges.CollegeCode)
		fmt.Println("Department -", i.Colleges.CollegeDepartment)
		fmt.Println()
		fmt.Println("=======================================")
	}

	return "All Data Caught"
}
func main() {
	f := FormDetails{
		Students: []StDetails{
			{
				Name: "Chandu_kumar",
				Roll: 12,
				Age:  23,
				Ad: Addres{
					Village:       "Adma",
					Post:          "Krashawa",
					policeStation: "Obra",
					PinCode:       824121,
					District:      "Aurangabad",
					State:         "Bihar",
				},
			},
			{
				Name: "Soanal",
				Roll: 23,
				Age:  28,
				Ad: Addres{
					Village:       "Bumaru",
					Post:          "Sundarganj",
					policeStation: "Amba",
					PinCode:       123422,
					District:      "Aurangabad",
					State:         "State",
				},
				Colleges: CollegeDetails{
					CollegeName:       "Sinha college Aurangabad",
					CollegeCode:       152,
					CollegeDepartment: "Computer Application",
				},
			},
			{
				Name: "Rambhu kumar",
				Roll: 43,
				Age:  23,
				Ad: Addres{
					Village:       "Adma",
					Post:          "Krashawa",
					policeStation: "Obra",
					PinCode:       824121,
					District:      "Aurangabad",
					State:         "Bihar",
				},
				Colleges: CollegeDetails{
					CollegeName:       "Yadav College Aurangabad",
					CollegeCode:       201,
					CollegeDepartment: "Math",
				},
			},
			{
				Name: "Ravi",
				Roll: 34,
				Age:  20,
				Ad: Addres{
					Village:       "Adma",
					Post:          "Krashawa",
					policeStation: "Obra",
					PinCode:       824121,
					District:      "Aurangabad",
					State:         "Bihar",
				},
				Colleges: CollegeDetails{
					CollegeName:       "Adarsh high School",
					CollegeCode:       203,
					CollegeDepartment: "Intermediate",
				},
			},
			{
				Name: "Randhir kumar",
				Roll: 89,
				Age:  19,
				Ad: Addres{
					Village:       "Adma",
					Post:          "Krashawa",
					policeStation: "Obra",
					PinCode:       824121,
					District:      "Aurangabad",
					State:         "Bihar",
				},
				Colleges: CollegeDetails{
					CollegeName:       "Sirish college Sirish",
					CollegeCode:       789,
					CollegeDepartment: "Intermediate",
				},
			},
			{
				Name: "Manish kumar",
				Roll: 90,
				Age:  29,
				Ad: Addres{
					Village:       "Basdiha",
					Post:          "Tengra",
					policeStation: "Barun",
					PinCode:       897654,
					District:      "Aurangabad",
					State:         "Bihar",
				},
				Colleges: CollegeDetails{
					CollegeName:       "Maisur",
					CollegeCode:       765,
					CollegeDepartment: "Computer Application",
				},
			},
			{
				Name: "paras kumar",
				Roll: 234,
				Age:  23,
				Ad: Addres{
					Village:       "Madanpur",
					Post:          "Rani kua",
					policeStation: "Madanpur",
					PinCode:       678659,
					District:      "Aurangabad",
					State:         "Bihar",
				},
				Colleges: CollegeDetails{
					CollegeName:       "Sinha college Aurangabad",
					CollegeCode:       152,
					CollegeDepartment: "Bsc.IT",
				},
			},
		},
	}
	fmt.Println(f.studentForm())

}
