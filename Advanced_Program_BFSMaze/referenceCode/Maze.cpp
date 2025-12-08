// Implementation of Maze class
// By Mary Elaine Califf and Jake Clouse

#include "Maze.h"
using namespace std;

#include "DisjointSet.h"

bool left(int randomCell, int numColumns, CellWalls mazeWalls[], DisjointSet &set);

bool down(int randomCell, int numColumns, CellWalls mazeWalls[], int mazeSize, DisjointSet &set);

bool right(int randomCell, int numColumns, CellWalls mazeWalls[], DisjointSet &set);

bool up(int randomCell, int mazeSize, int numColumns, CellWalls mazeWalls[], DisjointSet &set);


Maze::Maze(int rows, int cols, bool stop)
{
    numRows = rows;
    numColumns = cols;
    stopEarly = stop;
    int numCells = rows * cols;
    mazeWalls = new CellWalls[numCells];
    mazeWalls[numCells - 1].east = false;
}

Maze &Maze::operator=(const Maze &rhs)
{
    if (this != &rhs)
    {
        delete [] mazeWalls;
        this->copy(rhs);
    }
    return *this;
}

void Maze::generateMaze()
{
    int numCells = numRows * numColumns;
    DisjointSet mySet(numCells);
    bool mazeComplete = false;
    while(((!mazeComplete)&&(!stopEarly)) || (stopEarly && !(mySet.find(0) == mySet.find(numCells-1)))){
        int randomCell = rand() % numCells;
        int randomDirection = rand() % 4;
        //Left
        if(randomDirection == 0){
            mazeComplete = left(randomCell, numColumns, mazeWalls, mySet);
        }
        //Down
        else if(randomDirection == 1){
            mazeComplete = down(randomCell, numColumns, mazeWalls, numCells, mySet);
        }
        //Right
        else if(randomDirection == 2){
            mazeComplete = right(randomCell, numColumns, mazeWalls, mySet);
        }
        //UP
        else{
            mazeComplete = up(randomCell, numCells, numColumns, mazeWalls, mySet);
        }
    }
}

void Maze::print(ostream &outputStream)
{
    // print the top row of walls
    for (int i = 0; i < numColumns; i++)
        outputStream << " _";
    outputStream << '\n';
    for (int i = 0; i < numRows; i++)
    {
        int cellbase = i * numColumns;
        // print west wall (except at entrance)
        if (i == 0)
            outputStream << ' ';
        else
            outputStream << '|';
        for (int j = 0; j < numColumns; j++)
        {
            if (mazeWalls[cellbase + j].south)
                outputStream << '_';
            else
                outputStream << ' ';
            if (mazeWalls[cellbase + j].east)
                outputStream << '|';
            else
                outputStream << ' ';
        }
        outputStream << '\n';
    }
}

void Maze::copy(const Maze &orig)
{
    this->numRows = orig.numRows;
    this->numColumns = orig.numColumns;
    int numCells = numRows * numColumns;
    mazeWalls = new CellWalls[numCells];
    for (int i = 0; i < numCells; i++)
    {
        this->mazeWalls[i] = orig.mazeWalls[i];
    }
}

bool left(int randomCell, int numColumns, CellWalls mazeWalls[], DisjointSet &set){
    int neighborCell = 0;
    if(randomCell % numColumns == 0){
        return right(randomCell, numColumns, mazeWalls, set);
    }else{
        neighborCell = randomCell - 1;
        if(set.find(randomCell)!=set.find(neighborCell)){
            mazeWalls[neighborCell].east = false;
            return set.doUnion(randomCell, neighborCell);
        }
    }
    return false;
}

bool down(int randomCell, int numColumns, CellWalls mazeWalls[], int mazeSize, DisjointSet &set){
    int neighborCell = 0;
    if(randomCell + numColumns >= mazeSize){
        return up(randomCell, mazeSize, numColumns, mazeWalls, set);
    }else{
        neighborCell = randomCell + numColumns;
        if(set.find(randomCell)!=set.find(neighborCell)){
            mazeWalls[randomCell].south = false;
            return set.doUnion(randomCell, neighborCell);
        }
    }
    return false;

}

bool up(int randomCell, int mazeSize, int numColumns, CellWalls mazeWalls[], DisjointSet &set){
    int neighborCell = 0;
    if(randomCell - numColumns < 0){
        return down(randomCell, numColumns, mazeWalls, mazeSize, set);
    }else{
        neighborCell = randomCell - numColumns;
        if(set.find(randomCell)!=set.find(neighborCell)){
            mazeWalls[neighborCell].south = false;
            return set.doUnion(randomCell, neighborCell);
        }
    }
    return false;
}

bool right(int randomCell, int numColumns, CellWalls mazeWalls[], DisjointSet &set){
    int neighborCell = 0;
    if(randomCell % numColumns == numColumns-1){
        return left(randomCell, numColumns, mazeWalls, set);
    }else{
        neighborCell = randomCell + 1;
        if(set.find(randomCell)!=set.find(neighborCell)){
            mazeWalls[randomCell].east = false;
            return set.doUnion(randomCell, neighborCell);
        }
    }
    return false;
}


