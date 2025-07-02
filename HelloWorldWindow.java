import javax.swing.*;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;

public class HelloWorldWindow {

    public static void main(String[] args) {
createAndShowWindow();
    }
    public static void createAndShowWindow(){
        JFrame.setDefaultLookAndFeelDecorated(true);
        JFrame frame = new JFrame("Hola mundo Swing con FeHeJo");
JLabel text = new JLabel("Hola Mundo con FeHeJo");
//frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
frame.addWindowListener(new WindowAdapter() {
    @Override
    public void windowClosing(WindowEvent e) {
        System.out.println("cerrando ventana con exito");
        System.exit(0);
    }
});

frame.getContentPane().add(text);
frame.setVisible(true);
frame.pack();
    }
}
