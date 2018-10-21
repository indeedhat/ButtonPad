//
// Created by 4423 on 10/21/2018.
//

#ifndef BUTTON_PAD_MACROSTORE_H
#define BUTTON_PAD_MACROSTORE_H

#include <stdint-gcc.h>

class MacroStore
{
public:
  static const int BUTTONS = 16;

public:
  static bool save(int i, uint8_t data[]);
  static bool load(int i, uint8_t *data[]);

private:
  static bool check(int i, uint8_t data[]);
  static int maxSize();
};

#endif //BUTTON_PAD_MACROSTORE_H
