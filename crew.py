class CrewMember:
    def __init__(self, name, role):
        self.name = name
        self.role = role

    def introduce(self):
        return f"I be {self.name}, the {self.role}!"
