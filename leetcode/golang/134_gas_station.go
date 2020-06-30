func canCompleteCircuit(gas []int, cost []int) int {
    s, t, t1:= 0, 0, 0
    for i:= 0; i < len(gas); i++{
        t = t + gas[i] - cost[i]
        if (t < 0){
            s, t1, t = i + 1, t1 + t, 0
        }
    }
    if t1 + t < 0{
        return -1
    }
    return s
}
