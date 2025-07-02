import javax.swing.*;
import java.awt.*;

public class Window12ConfirmDialog extends JFrame {


Container panel;

public Window12ConfirmDialog()throws HeadlessException {

super("Confirmo antes de ejecutar una tarea fehejo");

panel =getContentPane();
 panel.setLayout(new FlowLayout());

 JPanel formPanel = new JPanel(new GridLayout(4,2,100,100));

 JLabel name = new JLabel("nombre: ", JLabel.RIGHT);
 JTextField nameField = new JTextField(5);
 formPanel.add(name);
 formPanel.add(nameField);

 JLabel lastname = new JLabel("apellido: ", JLabel.RIGHT);
 JTextField lastnameField = new JTextField();
 formPanel.add(lastname);
 formPanel.add(lastnameField);

 JLabel taxNumber = new JLabel("Rut: ", JLabel.RIGHT);
  JTextField taxNumberField = new JTextField();
 formPanel.add(taxNumber);
 formPanel.add(taxNumberField);

 JRadioButton morning = new JRadioButton("Grupo mañana", true);
 JRadioButton afternoon = new JRadioButton("Grupo tarde");

 formPanel.add(morning);
 formPanel.add(afternoon);

    int option = JOptionPane.showConfirmDialog(
            this,
            formPanel, "Introduce los datos",
            JOptionPane.OK_OPTION,
            JOptionPane.PLAIN_MESSAGE);
   System.out.println(option);

if (option==JOptionPane.OK_OPTION){
    System.out.println("Hemos selecionado OK");
    String nameValue = nameField.getText();
    System.out.println(nameValue);

    String lastnameValue = lastnameField.getText();
    System.out.println(lastnameValue);

    String taxNumberValue = taxNumberField.getText();
    System.out.println(taxNumberValue);

    JOptionPane.showMessageDialog(this, "Ejecutado con exito", "alerta", JOptionPane.INFORMATION_MESSAGE);


}else if (option== JOptionPane.CANCEL_OPTION){
    JOptionPane.showMessageDialog(this, "Ejecutado con error ", "alerta", JOptionPane.ERROR_MESSAGE);
        System.out.println("Hemos selecionado cancelar");

pack();
       // setSize(400,200);
        setVisible(true);
        setDefaultCloseOperation(EXIT_ON_CLOSE);
    }

}

    public static void main(String[] args) {
        new Window12ConfirmDialog();
    }
}


