// DisjointSet implementation using union by size and path compression
// By Mary Elaine Califf and Jake Clouse

#include "DisjointSet.h"
#include <iostream>

DisjointSet::DisjointSet(int numObjects){
    //to do
    numValues = numObjects;
    theArray = std::vector<int>(numObjects, -1);
}

//recursive method to find the item -- does path compression on the way out of the recursion
int DisjointSet::find(int objectIndex){
    // to do -- see assignment instructions for details
    if( theArray[ objectIndex ] < 0 )
        return objectIndex;
    else
        return theArray[ objectIndex ] = find( theArray[ objectIndex ] );
}

bool DisjointSet::doUnion(int objIndex1, int objIndex2){
    // to do -- see assignment instructions for details
    // printArrayValues(std::cout);

    bool foundNegative = false;
    int root1 = find(objIndex1);
    int root2 = find(objIndex2);
    if (theArray[root2]>=theArray[root1]){
       theArray[root1] += theArray[root2];
       theArray[root2] = root1;
    }
    else{
        theArray[root2] += theArray[root1];
        theArray[root1] = root2;
    }
    for(int i = 0; i < theArray.size(); i++){
        if(theArray[i] < 0){
            if(!foundNegative){
                foundNegative = true;
            }
            else{
                return false;
            }
        }
    }
    return true;
}

void DisjointSet::printArrayValues(std::ostream &outputStream){
    for (int i = 0; i < numValues; i++)
    {
        outputStream << theArray[i] << " ";
    }
    outputStream << std::endl;
}
