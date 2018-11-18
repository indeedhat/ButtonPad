#include <Arduino.h>
#include <Keyboard.h>
#include <PS2KeyAdvanced.h>
#include <SoftwareSerial.h>
#include "HID-Project.h"

// missing keys
#define KEY_MUTE     0x7F
#define KEY_VOL_UP   0x80
#define KEY_VOL_DOWN 0x81
#define KEY_0        0x27
#define KEY_1        0x1e
#define KEY_2        0x1f
#define KEY_3        0x20
#define KEY_4        0x21
#define KEY_5        0x22
#define KEY_6        0x23
#define KEY_7        0x24
#define KEY_8        0x25
#define KEY_9        0x26
#define KEY_DOT      0x37
#define KEY_EQL      0x67
#define KEY_MNS      0x56
#define KEY_PLS      0x57
#define KEY_DEV      0x54
#define KEY_MUL      0x55
#define KEY_K        0x0e

// settings
#define N_KEY_ROLLOVER 1


byte rows[]            = {10, 16, 14, 15, 18};
const uint8_t rowCount = 5;

byte cols[]            = {9, 8, 7, 6};
const uint8_t colCount = 4;


// layers
byte layers[]            = {5, 4, 3, 2};
const uint8_t layerCount = 4;
uint8_t layer            = 0;


// states
const uint8_t keyCount = 20;
bool stateChanged      = false;
byte state[keyCount];
bool modState[4];


// active key
uint8_t key = 0;


// setup the environment to be ready to read the key state
void setup() {
    for (uint8_t c = 0; c < colCount; c++) {
        pinMode(cols[c], INPUT);
    }

    for (uint8_t r = 0; r < rowCount; r++) {
        // pullup returns high if not connected to ground
        pinMode(rows[r], INPUT_PULLUP);
    }

    // setup the libs used for sending keystrokes
    Keyboard.begin();
    Consumer.begin();
    setLayerLED();
}


// read keys and set the state
void readKeys() {
    for (uint8_t c = 0; c < colCount; c++) {
        // send signal
        pinMode(cols[c], OUTPUT);
        digitalWrite(cols[c], LOW);

        for (uint8_t r = 0; c < rowCount; r++) {
            // get the key id
            key = c * rowCount + r;

            // read signal
            pinMode(rows[r], INPUT_PULLUP);

            // set key state
            byte tmp   = digitalRead(rows[r]);

            // mark as changed if needed
            stateChanged = stateChanged || byte != state[key];

            // set final state
            state[key] = tmp;

            // reset pin
            pinMode(rows[r], INPUT);
        }

        // disable signal for col
        pinMode(cols[c], INPUT);
    }
}


// ------------
// COMM HELPERS
// ------------

// initialize serial communication
bool setupSerial() {
    // dont do anything if serial is setup
    if (Serial) {
        return true;
    }

    Serial.begin(57600);
    delay(200);

    return !!Serial;
}

// send the current key and layer codes to the controller for processing
void sendToController() {
    if (!setupSerial()) {
        warningFlash();
        return;
    }

    // send key press info
    Serial.write(layer);
    Serial.write(key);

    // send message terminator
    Serial.write(";");
}


// -------------
// LAYER HELPERS
// -------------

// change the layer where required
#if 1 == N_KEY_ROLLOVER
void selectLayer() {
    uint8_t l = 0;
    for (uint8_t i = 0; i < keyCount; i += rowCount) {
        if (LOW == state[i]) {
            layer = l;
        }

        l++;
    }
}
#else
bool selectLayer() {
    uint8_t l = 0;

    for (uint8_t i = 0; i < keyCount; i += rowCount) {
        if (LOW == state[i]) {
            layer = l;
            return true;
        }

        l++;
    }

    return false;
}
#endif


// -----------
// LED HELPERS
// -----------

// flash all layer led's to let the user know that an error has happened
void warningFlash() {
    int power = HIGH;

    for (uint8_t i = 0; i < 10; i++) {
        pinMode(layers[0], 0 == layer ? HIGH : LOW);
        pinMode(layers[1], 1 == layer ? HIGH : LOW);
        pinMode(layers[2], 2 == layer ? HIGH : LOW);
        pinMode(layers[3], 3 == layer ? HIGH : LOW);
        delay(200);

        power = HIGH == power ? LOW : HIGH;
    }

    // return to the layer led display
    setLayerLED();
}

// change the led illumination for the correct active layer
void setLayerLED() {
    pinMode(layers[0], 0 == layer ? HIGH : LOW);
    pinMode(layers[1], 1 == layer ? HIGH : LOW);
    pinMode(layers[2], 2 == layer ? HIGH : LOW);
    pinMode(layers[3], 3 == layer ? HIGH : LOW);
}


// -----------------
// KEYSTROKE HELPERS
// -----------------

// execute keystrokes from a string of text
void text(char text[]) {
    Keyboard.print(text);
    delay(20);
}


// execute a sequence of key presses
void sequence(char keys[]) {
    for (int i = 0; i < sizeof(keys)/sizeof(keys[0]); i++) {
        // handle mod keys
        if (isMod(keys[i])) {
            toggleModState(keys[i]);

            if (getModState(keys[i])) {
                Keyboard.press(keys[i);
            } else {
                Keyboard.release(keys[i]);
            }

            continue;
        }

        Keyboard.press(keys[i]);
        delay(10);
        Keyboard.release(keys[i]);
    }

    Keyboard.releaseAll();
}


// check if a keycode is for a modifier key
bool isMod(char k) {
    return KEY_LEFT_SHIFT == k
        || KEY_LEFT_CTRL == k
        || KEY_LEFT_ALT == k
        || KEY_LEFT_GUI == k;
}


// check if a modifier key is currently being pressed
bool getModState(char k) {
    switch (k) {
        case KEY_LEFT_SHIFT:
            return modState[0];
        case KEY_LEFT_CTRL:
            return modState[1];
        case KEY_LEFT_ALT:
            return modState[2];
        case KEY_LEFT_GUI:
            return modState[3];
    }

    return false;
}


// change the state of a modifier key
void toggleModState(char k) {
    switch (k) {
        case KEY_LEFT_SHIFT:
            modState[0] = !modState[0];
            break;
        case KEY_LEFT_CTRL:
            modState[1] = !modState[1];
            break;
        case KEY_LEFT_ALT:
            modState[2] = !modState[2];
            break;
        case KEY_LEFT_GUI:
            modState[3] = !modState[3];
            break;
    }
}


// -----------------
// LAYER DEFINITIONS
// -----------------

// macro layers
// App open/web sites
void layer1(uint8_t i) {
    switch (i) {
        case 1: sendToController(); break; // IDEA
        case 2: sendToController(); break; // Spotify
        case 3: sendToController(); break; // MySQL Workbench
        case 4: sendToController(); break; // Meld

        case 6: sendToController(); break; // FileZilla
        case 7: sendToController(); break; // Thunder Bird
        case 8: sendToController(); break; // Chrome
        case 9: sendToController(); break; // Nvidia CP

        case 11: sendToController(); break; // Tiger VNC
        case 12: sendToController(); break; // Teamviewer
        case 13: break;
        case 14: break;

        case 16: text("http://github.com");     break;
        case 17: text("http://cloudflare.com"); break;
        case 18: text("http://10.0.0.113");     break;
        case 19: break;
    }
}


// Misc controls
void layer2(uint8_t i) {
    switch (i) {
        case 1: Consumer.send(MEDIA_PLAY_PAUSE);  break;
        case 2: Consumer.send(MEDIA_PREVIOUS);    break;
        case 3: Consumer.send(MEDIA_NEXT);        break;
        case 4: sequence((uint8_t[1]){KEY_MUTE}); break;

        case 6: break;
        case 7: break;
        case 8: sequence({KEY_VOL_DOWN}); break;
        case 9: sequence({KEY_VOL_UP});   break;

        case 11: break;
        case 12: break;
        case 13: break;
        case 14: break;

        case 16: sequence((uint8_t[3]){KEY_LEFT_CTRL, KEY_LEFT_ALT, KEY_DEL}); break;
        case 17: break;
        case 18: break;
        case 19: text("Fuck you Tim!"); break;
    }
}


// Shell/IDE Controls
void layer3(uint8_t i) {
    switch (i) {
        case 1: sequence((uint8_t[4]){KEY_LEFT_CTRL, KEY_LEFT_SHIFT, KEY_LEFT_ALT, KEY_K}); break;
        case 2: break;
        case 3: break;
        case 4: break;

        case 6: break;
        case 7: break;
        case 8: break;
        case 9: break;

        case 11: sequence({KEY_LEFT_CTRL, KEY_TAB})                                 break;
        case 12: text("git push origin master");                                    break;
        case 13: text("git commit -am ''"); sequence((uint8_t[1]){KEY_LEFT_ARROW}); break;
        case 14: sequence({KEY_RETURN});                                            break;

        case 16: text("cd ~");                               break;
        case 17: text("cd ~/Documents/GitHub");              break;
        case 18: text("cd ~/Documents/Project\\ Documents"); break;
        case 19: text("cd /var/www/vhosts");                 break;
    }
}


// key pad
void layer4(uint8_t i) {
    switch (i) {
        case 1: sequence((uint8_t[1]){KEY_7});   break;
        case 2: sequence((uint8_t[1]){KEY_8});   break;
        case 3: sequence((uint8_t[1]){KEY_9});   break;
        case 4: sequence((uint8_t[1]){KEY_MUL}); break;

        case 6: sequence((uint8_t[1]){KEY_4});   break;
        case 7: sequence((uint8_t[1]){KEY_5});   break;
        case 8: sequence((uint8_t[1]){KEY_6});   break;
        case 9: sequence((uint8_t[1]){KEY_DEV}); break;

        case 11: sequence((uint8_t[1]){KEY_1});   break;
        case 12: sequence((uint8_t[1]){KEY_2});   break;
        case 13: sequence((uint8_t[1]){KEY_3});   break;
        case 14: sequence((uint8_t[1]){KEY_PLS}); break;

        case 16: sequence((uint8_t[1]){KEY_DOT}); break;
        case 17: sequence((uint8_t[1]){KEY_0});   break;
        case 18: sequence((uint8_t[1]){KEY_EQL}); break;
        case 19: sequence((uint8_t[1]){KEY_MNS}); break;
    }
}


// ---------
// MAIN LOOP
// ---------

void loop() {
    readKeys();
#if 1 == N_KEY_ROLLOVER
    selectLayer();
#else
    if (selectLayer()) {
        return;
    }
#endif
    for (uint8_t i = 0; i < keyCount; i++) {
        // skip layer buttons
        if (0 == i || 0 == i % rowCount) {
            continue;
        }

        // skip keys that are not pressed
        if (LOW != state[i]) {
            continue;
        }

        // run the macros
        switch (layer) {
            case 0: layer1(i); break;
            case 1: layer2(i); break;
            case 2: layer3(i); break;
            case 3: layer4(i); break;
        }

#if 1 != N_KEY_ROLLOVER
        delay(10);
        Keyboard.releaseAll();
        return
#endif
    }

    delay(10);
    Keyboard.releaseAll();
}