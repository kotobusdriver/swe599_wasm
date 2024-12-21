"""
WASI Wall Clock is a clock API intended to let users query the current
time. The name "wall" makes an analogy to a "clock on the wall", which
is not necessarily monotonic as it may be reset.

It is intended to be portable at least between Unix-family platforms and
Windows.

A wall clock is a clock which measures the date and time according to
some external reference.

External references may be reset, so this clock is not necessarily
monotonic, making it unsuitable for measuring elapsed time.

It is intended for reporting the current date and time for humans.
"""
from typing import TypeVar, Generic, Union, Optional, Protocol, Tuple, List, Any, Self
from types import TracebackType
from enum import Flag, Enum, auto
from dataclasses import dataclass
from abc import abstractmethod
import weakref

from ..types import Result, Ok, Err, Some


@dataclass
class Datetime:
    """
    A time and date in seconds plus nanoseconds.
    """
    seconds: int
    nanoseconds: int


def now() -> Datetime:
    """
    Read the current value of the clock.
    
    This clock is not monotonic, therefore calling this function repeatedly
    will not necessarily produce a sequence of non-decreasing values.
    
    The returned timestamps represent the number of seconds since
    1970-01-01T00:00:00Z, also known as [POSIX's Seconds Since the Epoch],
    also known as [Unix Time].
    
    The nanoseconds field of the output is always less than 1000000000.
    
    [POSIX's Seconds Since the Epoch]: https://pubs.opengroup.org/onlinepubs/9699919799/xrat/V4_xbd_chap04.html#tag_21_04_16
    [Unix Time]: https://en.wikipedia.org/wiki/Unix_time
    """
    raise NotImplementedError

def resolution() -> Datetime:
    """
    Query the resolution of the clock.
    
    The nanoseconds field of the output is always less than 1000000000.
    """
    raise NotImplementedError
