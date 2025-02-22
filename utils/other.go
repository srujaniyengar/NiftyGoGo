package utils  
import "fmt"  
/** 
 * THIS MOD IS ONLY FOR CONVENIENT USE OF EXCEPTION HANDLING.  
 * THIS FUNCTION TAKES ANY ERROR AND PRINTS AN ERROR MESSAGE IF THE PROVIDED VALUE IS NOT NIL.  
 * THIS PREVENTS REDUNDANT ERROR CHECKING THROUGHOUT THE CODE.  
 */
func Check(err any) {
    if err != nil {
        fmt.Println(err, "Error")
        return
    }
}

