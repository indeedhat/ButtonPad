# CMAKE generated file: DO NOT EDIT!
# Generated by "MinGW Makefiles" Generator, CMake Version 3.8

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

SHELL = cmd.exe

# The CMake executable.
CMAKE_COMMAND = "C:\Program Files\JetBrains\CLion 2017.2\bin\cmake\bin\cmake.exe"

# The command to remove a file.
RM = "C:\Program Files\JetBrains\CLion 2017.2\bin\cmake\bin\cmake.exe" -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = C:\Users\4423\IdeaProjects\ButtonPad\Arduino\button_pad

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = C:\Users\4423\IdeaProjects\ButtonPad\Arduino\button_pad\cmake-build-debug

# Utility rule file for button_pad-size.

# Include the progress variables for this target.
include CMakeFiles/button_pad-size.dir/progress.make

CMakeFiles/button_pad-size: button_pad.elf
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --blue --bold --progress-dir=C:\Users\4423\IdeaProjects\ButtonPad\Arduino\button_pad\cmake-build-debug\CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Calculating button_pad image size"
	"C:\Program Files\JetBrains\CLion 2017.2\bin\cmake\bin\cmake.exe" -DFIRMWARE_IMAGE=C:/Users/4423/IdeaProjects/ButtonPad/Arduino/button_pad/cmake-build-debug/button_pad.elf -DMCU=atmega328p -DEEPROM_IMAGE=C:/Users/4423/IdeaProjects/ButtonPad/Arduino/button_pad/cmake-build-debug/button_pad.eep -P C:/Users/4423/IdeaProjects/ButtonPad/Arduino/button_pad/cmake-build-debug/CMakeFiles/FirmwareSize.cmake

button_pad-size: CMakeFiles/button_pad-size
button_pad-size: CMakeFiles/button_pad-size.dir/build.make

.PHONY : button_pad-size

# Rule to build all files generated by this target.
CMakeFiles/button_pad-size.dir/build: button_pad-size

.PHONY : CMakeFiles/button_pad-size.dir/build

CMakeFiles/button_pad-size.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles\button_pad-size.dir\cmake_clean.cmake
.PHONY : CMakeFiles/button_pad-size.dir/clean

CMakeFiles/button_pad-size.dir/depend:
	$(CMAKE_COMMAND) -E cmake_depends "MinGW Makefiles" C:\Users\4423\IdeaProjects\ButtonPad\Arduino\button_pad C:\Users\4423\IdeaProjects\ButtonPad\Arduino\button_pad C:\Users\4423\IdeaProjects\ButtonPad\Arduino\button_pad\cmake-build-debug C:\Users\4423\IdeaProjects\ButtonPad\Arduino\button_pad\cmake-build-debug C:\Users\4423\IdeaProjects\ButtonPad\Arduino\button_pad\cmake-build-debug\CMakeFiles\button_pad-size.dir\DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/button_pad-size.dir/depend
