hum samjhte hain kaise hum saare zeroes ko array ke end mein move kar sakte hain, aur non-zero elements ko unke relative order mein rakh sakte hain.

Problem Explanation:
Tumhare paas ek array hai: [0, 1, 0, 3, 12]
Tumhe saare zeroes ko end mein shift karna hai aur non-zero elements ko unke order mein rakhna hai.
Result aisa hona chahiye: [1, 3, 12, 0, 0]

Steps:
Ek index rakho jo track karega ki next non-zero element kahan jana chahiye.
Array ke har element ko dekho, agar wo zero nahi hai toh usse us index par rakho aur index ko aage badhao.
Array ke baaki positions ko zero se fill kar do.
Code:

go

package main

import "fmt"

func moveZeroes(nums []int) {
// Next position of the non-zero element
lastNonZeroFoundAt := 0

    // Move all the non-zero elements to the beginning of the array
    for _, num := range nums {
        if num != 0 {
            nums[lastNonZeroFoundAt] = num
            lastNonZeroFoundAt++
        }
    }

    // Fill remaining positions with zeroes
    for i := lastNonZeroFoundAt; i < len(nums); i++ {
        nums[i] = 0
    }

}

func main() {
nums := []int{0, 1, 0, 3, 12}
moveZeroes(nums)
fmt.Println(nums) // Output: [1, 3, 12, 0, 0]
}

Explanation:

1. Initialize lastNonZeroFoundAt:
   Yeh index track karega ki next non-zero element kahan jana chahiye.
   Initially, lastNonZeroFoundAt = 0.
2. Loop through the array:
   Har element ko dekho:
   Agar wo non-zero hai, toh usse lastNonZeroFoundAt position par rakho aur index ko increment karo.
3. Fill remaining positions with zeroes:
   Jab saare non-zero elements ko unke places pe rakh diya, toh baaki positions ko zero se fill kar do.

Example Ke Saath:
Array: [0, 1, 0, 3, 12]

Step Current Number (num) Array State lastNonZeroFoundAt
1 0 [0, 1, 0, 3, 12] 0
2 1 [1, 1, 0, 3, 12] 1
3 0 [1, 1, 0, 3, 12] 1
4 3 [1, 3, 0, 3, 12] 2
5 12 [1, 3, 12, 3, 12] 3
Ab lastNonZeroFoundAt = 3. Array ke remaining positions ko zero se fill kar do:

Final Array: [1, 3, 12, 0, 0]
I hope yeh samajh aaya! Agar koi doubt ho ya aur example chahiye, toh batao.

lekin ek or question humne fill karne ke liye loop ko lastNonZeroFoundAt se kyun chalaya wahan value nahi hai kya
ChatGPT
Bohot accha sawal, Salman! Chalo, yeh baat aur achhi tarah samajhte hain.

Fill Zeroes Loop:
Jab hum lastNonZeroFoundAt tak pohoch jate hain, toh iska matlab yeh hai ki humne saare non-zero elements ko sahi positions par rakh diya hai. Ab hume bas bache hue positions ko zero se fill karna hai.

Let's Break it Down:
Initial Loop:
Hum non-zero elements ko aage laate hain aur unhe sahi positions par rakhte hain.
Example with Explanation:
Suppose array hai: [0, 1, 0, 3, 12]

Loop 1: Move Non-Zero Elements

for \_, num := range nums {
if num != 0 {
nums[lastNonZeroFoundAt] = num
lastNonZeroFoundAt++
}
}
Initial State: [0, 1, 0, 3, 12], lastNonZeroFoundAt = 0
num = 0: Array same, lastNonZeroFoundAt = 0
num = 1: [1, 1, 0, 3, 12], lastNonZeroFoundAt = 1
num = 0: Array same, lastNonZeroFoundAt = 1
num = 3: [1, 3, 0, 3, 12], lastNonZeroFoundAt = 2
num = 12: [1, 3, 12, 3, 12], lastNonZeroFoundAt = 3
Ab array state hai [1, 3, 12, 3, 12] aur lastNonZeroFoundAt = 3.

Loop 2: Fill Remaining Positions with Zeroes
go
Copy code
for i := lastNonZeroFoundAt; i < len(nums); i++ {
nums[i] = 0
}
i = 3: [1, 3, 12, 0, 12]
i = 4: [1, 3, 12, 0, 0]
Why lastNonZeroFoundAt?
lastNonZeroFoundAt hume batata hai ki last non-zero element kahan rakha gaya tha.
Hum uske baad wale saare positions ko zero se fill karte hain.
Isliye loop ko lastNonZeroFoundAt se chalate hain taaki sirf un positions ko zero se fill karein jo already non-zero nahi hain.
Agar hum lastNonZeroFoundAt ke pehle se loop chalayenge, toh non-zero elements jo sahi jagah par hain, unhe zero se replace kar denge, jo galat hoga.
