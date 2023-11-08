//https://recursionist.io/dashboard/problems/524
class Assertion {
    public static void run(boolean b) throws Exception {
        if (!b) throw new Exception("Assertion Error");
        else System.out.println("Passed the assertion test.");
    }
}

class MyClass {
    public static int applyDiscount(int originalPrice, double discount) {
        double discountedPrice; 
        if (discount > 0 && discount < 1) {
            discountedPrice = originalPrice * (1.0 - discount);
        } else {
            discountedPrice = originalPrice;
        }
        return (int) discountedPrice;
    }

    public static void main(String[] args) throws Exception {
        // ここから書きましょう
        Assertion.run(applyDiscount(30,0.5) == 15);
        Assertion.run(applyDiscount(50,0.2) == 40);
        Assertion.run(applyDiscount(70,1.0) == 70);
        Assertion.run(applyDiscount(70,0.0) == 70);
        Assertion.run(applyDiscount(90,2.0) == 90);
        Assertion.run(applyDiscount(30,-3.0) == 30);

    }
}