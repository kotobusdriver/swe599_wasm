from typing import TypeVar, Generic, Union, Optional, Protocol, Tuple, List, Any, Self
from types import TracebackType
from enum import Flag, Enum, auto
from dataclasses import dataclass
from abc import abstractmethod
import weakref

from ..types import Result, Ok, Err, Some
from ..imports import network
from ..imports import udp


def create_udp_socket(address_family: network.IpAddressFamily) -> udp.UdpSocket:
    """
    Create a new UDP socket.
    
    Similar to `socket(AF_INET or AF_INET6, SOCK_DGRAM, IPPROTO_UDP)` in POSIX.
    On IPv6 sockets, IPV6_V6ONLY is enabled by default and can't be configured otherwise.
    
    This function does not require a network capability handle. This is considered to be safe because
    at time of creation, the socket is not bound to any `network` yet. Up to the moment `bind` is called,
    the socket is effectively an in-memory configuration object, unable to communicate with the outside world.
    
    All sockets are non-blocking. Use the wasi-poll interface to block on asynchronous operations.
    
    # Typical errors
    - `not-supported`:     The specified `address-family` is not supported. (EAFNOSUPPORT)
    - `new-socket-limit`:  The new socket resource could not be created because of a system limit. (EMFILE, ENFILE)
    
    # References:
    - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/socket.html>
    - <https://man7.org/linux/man-pages/man2/socket.2.html>
    - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-wsasocketw>
    - <https://man.freebsd.org/cgi/man.cgi?query=socket&sektion=2>
    
    Raises: `yana_app.types.Err(yana_app.imports.network.ErrorCode)`
    """
    raise NotImplementedError

