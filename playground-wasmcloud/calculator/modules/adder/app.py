from adder_component import exports

class Adder(exports.Adder):
    def add(self, a: int, b: int) -> int:
        return a + b