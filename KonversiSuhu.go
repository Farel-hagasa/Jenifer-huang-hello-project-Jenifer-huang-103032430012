package main
// menambahkan komentar
import "fmt"

func Reamur(C float64) float64{

	return (4.0/5.0)*C
}	
func Fahrenheit(C float64) float64{
	
	return  (9.0/5.0 ) * C + 32
}	
func Kelvin(C float64) float64{
	
	return  (C + 273)
	
}

func main() {
	var awal, akhir , step float64
	var  C , R , F ,K float64

	fmt.Scan(&awal,&akhir,&step)
	for C = awal; C<= akhir; C+= step {
		R = Reamur(C)
		F = Fahrenheit(C)
		K = Kelvin(C)
		fmt.Println("\t\t\t\t\n",C,R,F,K)
	}	
		
}