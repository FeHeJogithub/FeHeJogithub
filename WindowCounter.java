import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class WindowCounter extends JFrame implements ActionListener {


    private int counter;
    private JButton button = new JButton("Count");
    private JLabel result = new JLabel("Counter");

    private JButton buttonDecrement = new JButton("Decrement");


    public WindowCounter() throws HeadlessException {
        super("Ejemplo contador con fehejo");

        Container panel = getContentPane();
        panel.add(button);

        panel.add(new JLabel("-"));


        panel.add(result);
        panel.setLayout(new FlowLayout());
        button.addActionListener(this);
        panel.add(buttonDecrement);
        buttonDecrement.addActionListener(event ->{
            counter--;
            result.setText("counter: "+ counter);
        });

        setVisible(true);
        setSize(400,100);
        setDefaultCloseOperation(EXIT_ON_CLOSE);

    }


    public static void main(String[] args) {
new WindowCounter();
    }

    @Override
    public void actionPerformed(ActionEvent e) {
        counter++;
        result.setText("Counter : ".concat(String.valueOf(counter)));
    }
}
