import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Window07CurrencyConverter extends JFrame implements ActionListener {

    Container panel;
    JButton buttonEu, buttonUsd, buttonRest;
    JTextField mount;

    public static final String CURRENCY_DOLAR= "Dolar";
    public static final String CURRENCY_EUROS= "Euros";

    public Window07CurrencyConverter() throws HeadlessException {
        super("Ejemplo de conversor de monedas con fehejo");

        panel = getContentPane();
        panel.setLayout(new FlowLayout());

        mount = new JTextField(10);
buttonRest = new JButton("Rest");
buttonRest.addActionListener(this);
        buttonEu = new JButton(CURRENCY_EUROS);


        buttonEu.addActionListener(this);
        buttonUsd = new JButton(CURRENCY_DOLAR);
        buttonUsd.addActionListener(this);

        panel.add(mount);
        panel.add(buttonEu);
        panel.add(buttonUsd);
panel.add(buttonRest);
        setSize(300, 250);
        setVisible(true);
        setDefaultCloseOperation(EXIT_ON_CLOSE);

    }

    @Override
    public void actionPerformed(ActionEvent e) {
        float value = Float.parseFloat(mount.getText());
        String textCurrency = e.getActionCommand();
        if (textCurrency.equals(CURRENCY_EUROS)){
            value = (value/66.19F);
        panel.setBackground(Color.green);

        }else if (textCurrency.equals(CURRENCY_DOLAR)){
            value = value /60.38F;

            panel.setBackground(Color.blue);
        }else{
            value = 0.0F;
        }

        mount.setText(String.valueOf(value));
        }




    public static void main(String[] args) {
        new Window07CurrencyConverter();
    }
}
