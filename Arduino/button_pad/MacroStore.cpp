//
// Created by 4423 on 10/21/2018.
//
#include "MacroStore.h"
#include <C:\Program Files (x86)\Arduino\hardware\arduino\avr\libraries\EEPROM\src\EEPROM.h>


static bool MacroStore::load(int i, uint8_t *data[])
{
  if (0 > i || MacroStore::BUTTONS <= i) {
    return false;
  }

  for (int x = 0, m = MacroStore::maxSize(); i < m; i++) {
    *data[x] = EEPROM.read(x);
  }

  return true;
}


static bool MacroStore::save(int i, uint8_t data[])
{
  if (!MacroStore::check(i, data)) {
    return false;
  }

  for (int x = 0; x < sizeof(data); x++) {
    EEPROM.write(i + x, data[x]);
  }

  return true;
}


static int MacroStore::maxSize()
{
  int total = EEPROM.length();
  return total / MacroStore::BUTTONS;
}


static bool MacroStore::check(int i, uint8_t data[])
{
  return 0 <= i
         && MacroStore::BUTTONS > i
         && sizeof(data) <= MacroStore::maxSize();
}