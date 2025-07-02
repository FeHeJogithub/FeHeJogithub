import javax.swing.*;
import javax.swing.table.AbstractTableModel;
import javax.swing.table.TableModel;

public class Window13TableModel extends JFrame {

    public Window13TableModel() {
        super("Tabla con Registro o Grilla fehejo");

        JTable table = new JTable(new UserTableModel());
        JScrollPane scroll = new JScrollPane(table);

        JPanel panel = new JPanel();
        panel.add(scroll);
        getContentPane().add(panel);

        setVisible(true);
        pack();
        setDefaultCloseOperation(EXIT_ON_CLOSE);

    }

    public static void main(String[] args) {
        new Window13TableModel();
    }

    private class UserTableModel extends AbstractTableModel {
        private String[] columns = {"Id", "Name", "LastName", "Email"};
        private Object[][] rows;

        public UserTableModel(){
            this.rows  =new Object[5][4];
        rows [0] = new Object[]{1, "Andres", "Guzman", "aguzman@gmail.com"};
        rows [1] = new Object[]{2, "Pepe", "Polanco", "polanco12@gmail.com"};
        rows [2] = new Object[]{3, "Ana", "Gomez", "agomez@gmail.com"};

        rows [3] = new Object[]{4, "James", "Rosado", "jrosado@gmail.com"};
        rows [4] = new Object[]{5, "Bruce", "Lee", "blee@gmail.com"};

        }



        @Override
        public int getRowCount() {
            return rows.length;
        }

        @Override
        public int getColumnCount() {
            return columns.length;
        }

        @Override
        public Object getValueAt(int rowIndex, int columnIndex) {
            return rows[rowIndex][columnIndex];
        }

        @Override
        public String getColumnName(int column) {
            return columns[column];
        }

    }
}

