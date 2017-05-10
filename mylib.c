#include <stdio.h>
#include "mylib.h"

// force gcc to link in go runtime (may be a better solution than this)
void dummy() {
    GoString gs;
    PrintGoStr(gs);
    PrintInt(123);
    PrintCStr("asd");
    Coords cds;
    PrintStruct(cds);
}

int main() {

}