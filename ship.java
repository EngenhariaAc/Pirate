public class Ship {
    private String name;
    private int cannons;

    public Ship(String name, int cannons) {
        this.name = name;
        this.cannons = cannons;
    }

    public void fireCannons() {
        System.out.println(name + " fires " + cannons + " cannons! Boom!");
    }
}
