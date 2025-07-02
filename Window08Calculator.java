import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Window08Calculator extends JFrame implements ActionListener {

    //Container panel;
    Container container;

    JTextField numberA, numberB, result;

    public Window08Calculator() throws HeadlessException {
        super("Calculadora con fehejo");

        container = getContentPane();
        container.setLayout(new BorderLayout());

        JMenuBar bar = new JMenuBar();
        setJMenuBar(bar);
        JMenu menu = new JMenu("Operaciones");

        JMenuItem itemAdd = new JMenuItem("Suma");
        JMenuItem itemSubt = new JMenuItem("Resta");
        JMenuItem itemDiv = new JMenuItem("Divide");
        JMenuItem itemMult = new JMenuItem("Multiplica");

        menu.add(itemAdd);
        menu.add(itemSubt);
        menu.add(itemDiv);
        menu.add(itemMult);

        JMenu closeMenu = new JMenu("Cerrar");
        JMenuItem itemClose = new JMenuItem("Salir");
        closeMenu.add(itemClose);

        bar.add(menu);
        bar.add(closeMenu);

        JPanel panel = new JPanel();
        panel.setLayout(new FlowLayout());
        panel.add(new JLabel("Numero 1"));
        panel.add(numberA = new JTextField(3));
        panel.add(new JLabel("Numero 2"));
        panel.add(numberB = new JTextField(3));
        panel.add(new JLabel("Resultado"));
        panel.add(result = new JTextField(5));
        result.setEditable(false);

        // 👉 ESTA LÍNEA AÑADE EL PANEL A LA VENTANA
        container.add(panel, BorderLayout.CENTER);

        itemAdd.addActionListener(this);
        itemSubt.addActionListener(this);
        itemDiv.addActionListener(this);
        itemMult.addActionListener(this);
        itemClose.addActionListener(this);

        pack(); // Acomoda el tamaño al contenido
        setVisible(true);
        setDefaultCloseOperation(EXIT_ON_CLOSE);
    }


    @Override
    public void actionPerformed(ActionEvent e) {

        String operation = e.getActionCommand();
        int a = Integer.parseInt(numberA.getText().trim());
        int b = Integer.parseInt(numberB.getText().trim());
        int value=0;


        switch (operation){
            case "Suma":
                value= a+b;
                break;
            case "Resta":
                value= a-b;
                break;
            case "Divide":
                if(b==0){throw new ArithmeticException("no se puede dividir en cero ");
            }
            value= a/b;
                break;
            case "Multiplica":
                value= a*b;
                break;

            case "Salir":
                System.exit(0);
        }
        result.setText(String.valueOf(value));
    }

    public static void main(String[] args) {
        new Window08Calculator();
}
}
