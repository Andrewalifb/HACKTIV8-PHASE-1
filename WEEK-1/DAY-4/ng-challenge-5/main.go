package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name string
	Age  int
	Job  string
}

func (p Person) GetInfo() {
	fmt.Printf("Hallo nama saya %s umur saya %d dan saya bekerja sebagai %s\n", p.Name, p.Age, p.Job)
}

func (y *Person) Addyear(year int) {
	y.Age += year
	if y.Age >= 50 {
		y.Job = "Retired"
		fmt.Printf("Umur %s saat ini adalah %d dan status pekerjaan %s\n",y.Name, y.Age, y.Job)
	} else {
		fmt.Printf("Umur %s saat ini adalah %d dan status pekerjaan %s\n",y.Name, y.Age, y.Job)
	}
}

func introEachPerson(data []Person) {
	for index, _ := range data {
		data[index].GetInfo()
	}
}

type Weapon struct {
  Attack int
}

type Hero struct {
  Name           string
  BaseAttack     int
  Defense        int
  CriticalDamage int
  HealthPoint    int
  Weapon         Weapon 
}

func (h *Hero) CountDamage() int {
	damage := rand.NewSource(time.Now().UnixNano())
	h.CriticalDamage = rand.New(damage).Intn(2)
	total_damage := h.BaseAttack + h.Weapon.Attack + h.CriticalDamage
	// fmt.Println("Total Damage :", total_damage)
return total_damage
}

func (hero *Hero) isAttackedBy(damage int) {
	// Jika damage lebih kecil atau sama dengan defense maka tidak akan mengurangi healt point
	if damage <= hero.Defense {
		armor := (hero.Defense - damage)
		if armor < 0 {
			hero.Defense = 0
		} else {
			hero.Defense -= damage
		}
	// Jika damage lebih besar dari defense maka akan mengurangi health point
	} else if damage > hero.Defense {
			remaining_damage := damage - hero.Defense
			hero.HealthPoint -= remaining_damage
			hero.Defense = 0
	}

}

func Battle(attackers, attacked Hero) {
	damage := attackers.CountDamage()
	attacked.isAttackedBy(damage)
	fmt.Println("*****BATTLE RESULT*****")
	fmt.Println("=======================")
	fmt.Println("Attackers : ", attackers.Name)
	fmt.Println("Attackers Attack Damage : ", damage)
	fmt.Printf("Hero Attacked: %s\n", attacked.Name)
	fmt.Printf("Get Damage : %d\n", damage)
	fmt.Printf("Remaning Defense : %d\n", attacked.Defense)
	fmt.Printf("Remaining Helath Point : %d\n", attacked.HealthPoint)
	fmt.Println("")
}

func main() {
	/* NG Challenge 5 : Struct & Method 1 */
  person1 := Person{Name : "Bambang", Age: 20, Job: "Gambler"}
  person2 := Person{Name : "Budi", Age: 45, Job: "Banker"}
	person1.GetInfo()
	person2.Addyear(1)
	person2.Addyear(1)
	person2.Addyear(1)
	person2.Addyear(1)
	person2.Addyear(1)

	/* NG Challenge 5 : Struct & Method 2 */
	anotherPerson := []Person{
		{Name: "Toni", Age: 20, Job: "Front End Engineer"},
		{Name: "Santi", Age: 25, Job: "Backend End Engineer"},
		{Name: "Laras", Age: 23, Job: "Full Stack Engineer"},
	}

	introEachPerson(anotherPerson)

	/* NG Challenge 5 : Struct HERO 1 */
  Marvel := Hero {
		Name:           "Superman",
		BaseAttack:     100,
		Defense:        50,
		CriticalDamage: 50,
		HealthPoint:    150,
		Weapon:         Weapon{Attack: 30},
	}


	DC := Hero {
		Name:           "Batman",
		BaseAttack:     80,
		Defense:        70,
		CriticalDamage: 30,
		HealthPoint:    180,
		Weapon:         Weapon{Attack: 25},
	}
fmt.Println("Iron Man Damage : ", Marvel.CountDamage())
fmt.Println("Batman Damage :",DC.CountDamage())
/* NG Challenge 5 : Struct HERO 1 */
Battle(Marvel, DC)
Battle(DC, Marvel)
}
