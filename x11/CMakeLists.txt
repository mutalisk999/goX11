# cmake -G "MinGW Makefiles" . && make

cmake_minimum_required(VERSION 3.10)

add_definitions()

set(X11_SRC
    aes_helper.cpp
    blake.cpp
    bmw.cpp
    cubehash.cpp
    echo.cpp
    groestl.cpp
    jh.cpp
    keccak.cpp
    luffa.cpp
    shavite.cpp
    simd.cpp
    skein.cpp
    x11.cpp
)

add_library(x11 STATIC ${X11_SRC})
target_link_libraries(x11)

