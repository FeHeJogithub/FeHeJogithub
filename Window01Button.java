import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Window01Button extends JFrame implements ActionListener {

    public Window01Button(){
        super("Ejemplo de Boton y y Eventos con FeHeJo");

        JButton button = new JButton("Aceptar");

        button.addActionListener(this);
        button.addActionListener(new ButtonClickListener());

        getContentPane().add(button);
        setSize(200,100);
        setVisible(true);
        setDefaultCloseOperation(EXIT_ON_CLOSE);

    }
    public static void main(String[] args) {
        new Window01Button();
    }

    @Override
    public void actionPerformed(ActionEvent e) {
        System.out.println("Boton pulsado con fehejo");
        System.out.println("Boton pulsado 2 con fehejo");
        Toolkit.getDefaultToolkit().beep();
    }

     static class ButtonClickListener implements ActionListener {
        @Override
        public void actionPerformed(ActionEvent e) {
            System.out.println("Boton pulsado 3 con fehejo");
        }
    }
}
