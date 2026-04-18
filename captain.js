class Captain {
  constructor(name, ship) {
    this.name = name;
    this.ship = ship;
  }

  speak() {
    console.log(`Ahoy! I be Captain ${this.name} of the ${this.ship}!`);
  }
}

module.exports = Captain;
