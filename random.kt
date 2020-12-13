/**
 * You can edit, run, and share this code. 
 * play.kotlinlang.org 
 */
import kotlin.random.Random
import java.time.LocalDate



fun main() {
   
   
   
      
    pinChallengeIndices.forEach{
        println(it)
    }
    
    println("--------------------------------------------------")
    passowrdChallengeIndices.forEach{
        println(it)
    }
    println("--------------------------------------------------")
    val saltAvailableIndices = ('A'..'Z').union('0'..'9').toList()
    
    val saltRes = saltIndices.map{
        saltAvailableIndices[it]
    }.joinToString("")        
    
    
    println(saltRes)
    println("--------------------------------------------------")   


}

fun getSaltIndices(value: Long, indiceMaxValue: Int, indiceMaxCount: Int):List<Int> {
    
    val longToCharArray = value.toString().toCharArray()
    val indiceMaxValueIncludingLastIndex = indiceMaxValue + 1
    val bumpFactor = 20
    val returnIndices  = mutableListOf<MutableList<Int>>()
    val mintedIndices = longToCharArray.map{
         Integer.valueOf(it.toString()) + bumpFactor % indiceMaxValueIncludingLastIndex
    }    
    return mintedIndices.take(indiceMaxCount)
}

fun getChallengeIndices(value: Long, indiceMaxValue: Int, indiceMaxCount: Int):List<Int> {
    val longToCharArray = value.toString().toCharArray()    
    val returnIndices: MutableList<Int> = mutableListOf()
    val indiceMaxValueIncludingLastIndex = indiceMaxValue + 1
    var index = 0

    longToCharArray.forEach{      
        if (returnIndices.size < indiceMaxCount){
            val mintedValue = Integer.valueOf(it.toString())% indiceMaxValueIncludingLastIndex
            if  (!returnIndices.contains(mintedValue) &&  mintedValue != 0 ){
                returnIndices.add(mintedValue)
                index++                
            }
        } 
    }
    
    if (returnIndices.size < indiceMaxCount){
        val missingIndicesLength = indiceMaxCount - returnIndices.size
        val potentialIndices = listOf(1..indiceMaxValue)
        val remainingIndices = potentialIndices.minus(returnIndices)
        
        for (i in 1 .. missingIndicesLength){
         	returnIndices.add(remainingIndices.get(i) as Int)   
        }
    }    
	return returnIndices
    
}
  
