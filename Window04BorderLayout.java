import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Window04BorderLayout extends JFrame implements ActionListener{

        public Window04BorderLayout(){
            super("Ejemplo de Boton y y Eventos con FeHeJo");


            getContentPane().setLayout (new BorderLayout(8, 4));

            //for (int i = 1; i <= 10; i ++){

                JButton button = new JButton("Aceptar");

                button.addActionListener(this);
                button.addActionListener(new ButtonClickListener());

                button.setPreferredSize(new Dimension(200,100));

                getContentPane().add(button, BorderLayout.CENTER);
                getContentPane().add(new JButton("North"), BorderLayout.NORTH);
                getContentPane().add(new JButton("South"), BorderLayout.SOUTH);
                getContentPane().add(new JButton("West"), BorderLayout.WEST);
                getContentPane().add(new JButton("East"), BorderLayout.EAST);


           // }


            setSize(600,200);

            setVisible(true);
            setDefaultCloseOperation(EXIT_ON_CLOSE);

        }
        public static void main(String[] args) {
            new Window04BorderLayout();
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


