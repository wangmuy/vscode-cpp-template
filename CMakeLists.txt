cmake_minimum_required(VERSION 3.12)
project(vscodecpptemplate)

set(TARGET_NAME ${CMAKE_PROJECT_NAME})

add_executable(${TARGET_NAME})
target_sources(${TARGET_NAME} PRIVATE
    main.cpp
)

include(CTest)
enable_testing()

set(CPACK_PROJECT_NAME ${PROJECT_NAME})
set(CPACK_PROJECT_VERSION ${PROJECT_VERSION})
include(CPack)
