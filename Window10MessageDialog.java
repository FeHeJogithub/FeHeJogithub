import javax.swing.*;
import java.awt.*;

public class Window10MessageDialog extends JFrame {

    Container p;

    public Window10MessageDialog() throws HeadlessException {
        super("Ventana de mensaje y alerta");

        p = getContentPane();
        p.setLayout(new FlowLayout());

        JOptionPane.showMessageDialog(null, "Debe introducir " +
                        "datos en los campos",
                "Error de entrada",JOptionPane.ERROR_MESSAGE);


        setSize(400,250);
        setVisible(true);
        setDefaultCloseOperation(EXIT_ON_CLOSE);
    }

    public static void main(String[] args) {
        new Window10MessageDialog();
    }



}
