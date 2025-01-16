from divider_component import exports

class Divider(exports.Divider):
    def divide(self, a: int, b: int) -> int:
        return a // b