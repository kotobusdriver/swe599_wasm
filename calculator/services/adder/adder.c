#include <emscripten.h>

EMSCRIPTEN_KEEPALIVE
int process(int arg1, int arg2){
  return arg1 + arg2;
}