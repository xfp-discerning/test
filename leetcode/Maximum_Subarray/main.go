package main

func maxSubArray(nums []int) int {
    var cursum, presum,maxsum int
    length := len(nums)
    for i:=1;i<=length;i++ {
        // if nums[i]< 0{
        //     continue
        // }
        cursum = presum + nums[i]
        if cursum>maxsum{
            maxsum = cursum
        }
        presum = maxsum
    }
    return maxsum
}