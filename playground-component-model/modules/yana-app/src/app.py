from yana_app import exports, multiply
from yana_app.imports import adder

class Run(exports.Run):
    def run(self) -> None:
        addition = adder.add(10, 20)
        multiplication = multiply(addition, 2)
        print(multiplication)