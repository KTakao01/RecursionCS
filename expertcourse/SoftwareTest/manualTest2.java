class MyClass {
    public static char getStudentGPA(int points) {
        char output = ' ';
        if (points < 0) output = 'F';
        if (points >= 0 && points < 50) output = 'D';
        if (points >= 50 && points <= 70) output = 'C';
        if (points >= 70 && points <= 80) output = 'B';
        if (points > 80) output = 'A';

        return output;
    }

    public static void unitTestCheck(boolean predicate) {
        if (predicate) {
            System.out.println("The test passed!!");
        } else {
            System.out.println("ERROR! The test failed!!");
        }
    }

    public static void unitTests() {
        // ここから書いてください
        unitTestCheck(getStudentGPA(-1) == 'F');
        unitTestCheck(getStudentGPA(0) == 'D');
        unitTestCheck(getStudentGPA(50) == 'C');
        unitTestCheck(getStudentGPA(69) == 'C');
        unitTestCheck(getStudentGPA(80) == 'A');
    }
}
