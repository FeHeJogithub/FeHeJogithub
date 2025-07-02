package org.fhernandez.java.jdbc;

import javax.swing.*;
import javax.swing.table.AbstractTableModel;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.util.ArrayList;
import java.util.List;


public class JdbcSwingCrudProjectModificacionValidacion extends JFrame {

private Container p;

private JTextField nameField = new JTextField();
private JTextField priceField = new JTextField();
private JTextField quantityField = new JTextField();
  //  private JTextField contryField = new JTextField();
private ProductTableModel tableModel = new ProductTableModel();

    public JdbcSwingCrudProjectModificacionValidacion() {
super("Swing: GUI con base de Datos Mysql");

p = getContentPane();
p.setLayout (new BorderLayout(20,10));

JPanel formPanel = new JPanel(new GridLayout(5,2,20,10));
formPanel.setBorder(BorderFactory.createEmptyBorder(20,20,20,20));

JButton buttonSave = new JButton("Guardar");

formPanel.add(new JLabel("Nombre: "));
formPanel.add(nameField);

formPanel.add(new JLabel("Price: "));
formPanel.add(priceField);

formPanel.add(new JLabel("Cantidad: "));
formPanel.add(quantityField);

//formPanel.add(new JLabel("Pais: "));
//formPanel.add(contryField);

formPanel.add(new JLabel(""));
formPanel.add(buttonSave);
buttonSave.addActionListener(new AddActionListener());

JPanel tablePanel = new JPanel(new FlowLayout());

JTable jTable = new JTable();
jTable.setModel(this.tableModel);

JScrollPane scroll = new JScrollPane(jTable);
tablePanel.add(scroll);

p.add(tablePanel, BorderLayout.SOUTH);

        p.add(formPanel, BorderLayout.NORTH);
        pack();
        setVisible(true);
        setDefaultCloseOperation(EXIT_ON_CLOSE);

    }

    public static void main(String[] args) {
new JdbcSwingCrudProjectModificacionValidacion();
    }

    private class AddActionListener implements ActionListener {
        @Override
        public void actionPerformed(ActionEvent e) {
           String name = nameField.getText();
           int price = Integer.parseInt(priceField.getText());
           int quantity = Integer.parseInt(quantityField.getText());
           //String contry = contryField.getText();

           Object[] product = new Object[]{System.currentTimeMillis(),name, price, quantity};
           tableModel.getRows().add(product);//crea un producto
           tableModel.fireTableDataChanged();

           nameField.setText("");
           priceField.setText("");
           quantityField.setText("");

          // System.out.println(product);
           System.out.println(product[0]);
           System.out.println(product[1]);
           System.out.println(product[2]);
         // System.out.println(product[3]);

        }
    }

    private class ProductTableModel extends AbstractTableModel {
        private String[] columns = new String[]{"Id", "Nombre", "Precio", "Cantidad"};
        private List<Object[]> rows =new ArrayList<>();

        public List<Object[]> getRows() {
            return rows;
        }

        @Override

        public int getRowCount() {
            return rows.size();
        }

        @Override
        public int getColumnCount() {
            return columns.length;
        }

        @Override
        public Object getValueAt(int rowIndex, int columnIndex) {
            return rows.get(rowIndex)[columnIndex];
        }

        @Override
        public String getColumnName(int column) {
            return columns[column];
        }
    }
}
