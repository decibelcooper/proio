/**
 * <h1> ProIO Browser</h1>
 * <p>
 * This program creates an event viewer and file browser. It allows a user to
 * select a file and display the number of events to choose. Then the user is
 * able to go to an event and gather the tags from that event. when a user
 * selects a tag it populates a table. In the table if the user selects an entry
 * they can then view its entry data.
 * <p>
 */


package proio;

import java.io.File;
import java.io.IOException;
import java.util.logging.Level;
import java.util.logging.Logger;
import javax.swing.DefaultListModel;
import javax.swing.JFileChooser;

import java.util.Map;

import com.google.protobuf.ByteString;
import com.google.protobuf.Message;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import javax.swing.JFrame;
import javax.swing.JOptionPane;
import javax.swing.JScrollPane;
import javax.swing.JTable;
import javax.swing.JTextArea;
import javax.swing.table.DefaultTableModel;
import org.apache.commons.lang3.StringUtils;
/**
 *
 * @author Jose de Jesus Alcaraz Jr.
 * @version 1.0
 * @since 8/08/18
 */
public class ProIOBrowser extends javax.swing.JFrame {

    /**
     * Creates new form proio_browser
     */
    public ProIOBrowser() {
        initComponents();
    }

    @SuppressWarnings("unchecked")                       
    private void initComponents() {

        DataPanel = new javax.swing.JPanel();
        jSplitPane1 = new javax.swing.JSplitPane();
        jSplitPane2 = new javax.swing.JSplitPane();
        jScrollPane1 = new javax.swing.JScrollPane();
        EntryTable = new javax.swing.JTable();
        jScrollPane5 = new javax.swing.JScrollPane();
        EntryData = new javax.swing.JTextArea();
        jScrollPane2 = new javax.swing.JScrollPane();
        TagList = new javax.swing.JList<String>();
        jPanel2 = new javax.swing.JPanel();
        File_label = new javax.swing.JLabel();
        eventnumbersearchfield = new javax.swing.JFormattedTextField();
        eventlabel = new javax.swing.JLabel();
        jPanel4 = new javax.swing.JPanel();
        PrevButton = new javax.swing.JButton();
        NextButton = new javax.swing.JButton();
        numbereventslabel = new javax.swing.JLabel();
        eventnumlabel = new javax.swing.JLabel();
        entryidtext = new javax.swing.JTextField();
        EnterButton = new javax.swing.JButton();
        jPanel3 = new javax.swing.JPanel();
        entriesLabel = new java.awt.Label();
        tagsLabel = new java.awt.Label();
        entrydatalabel = new java.awt.Label();
        jMenuBar1 = new javax.swing.JMenuBar();
        File_Menu = new javax.swing.JMenu();
        File_button = new javax.swing.JRadioButtonMenuItem();
        Meta_menu = new javax.swing.JMenu();
        HeaderRecord_button = new javax.swing.JRadioButtonMenuItem();
        ParticalData_button = new javax.swing.JRadioButtonMenuItem();
        Statistics_button = new javax.swing.JRadioButtonMenuItem();
        Logfile_button = new javax.swing.JRadioButtonMenuItem();
        DataLayout_menu = new javax.swing.JMenu();
        DesLay_button = new javax.swing.JRadioButtonMenuItem();
        HeadLay_button = new javax.swing.JRadioButtonMenuItem();
        EventLay_button = new javax.swing.JRadioButtonMenuItem();
        StatLay_button = new javax.swing.JRadioButtonMenuItem();
        Help_menu = new javax.swing.JMenu();
        About_button = new javax.swing.JRadioButtonMenuItem();

        setDefaultCloseOperation(javax.swing.WindowConstants.EXIT_ON_CLOSE);
        setTitle("Proio");
        setName("");

        DataPanel.setAutoscrolls(true);

        jSplitPane1.setAutoscrolls(true);

        jSplitPane2.setResizeWeight(1.0);
        jSplitPane2.setAutoscrolls(true);
        jSplitPane2.setContinuousLayout(true);

        EntryTable.setAutoCreateRowSorter(true);
        EntryTable.setModel(new javax.swing.table.DefaultTableModel(
            new Object [][] {

            },
            new String [] {
                "ID", "Message Type"
            }
        ) {
            Class[] types = new Class [] {
                java.lang.Long.class, java.lang.String.class
            };
            boolean[] canEdit = new boolean [] {
                false, false
            };

            public Class getColumnClass(int columnIndex) {
                return types [columnIndex];
            }

            public boolean isCellEditable(int rowIndex, int columnIndex) {
                return canEdit [columnIndex];
            }
        });
        EntryTable.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                EntryTableMouseClicked(evt);
            }
        });
        jScrollPane1.setViewportView(EntryTable);
        if (EntryTable.getColumnModel().getColumnCount() > 0) {
            EntryTable.getColumnModel().getColumn(0).setPreferredWidth(50);
            EntryTable.getColumnModel().getColumn(0).setMaxWidth(200);
        }

        jSplitPane2.setLeftComponent(jScrollPane1);

        EntryData.setColumns(20);
        EntryData.setRows(5);
        jScrollPane5.setViewportView(EntryData);

        jSplitPane2.setRightComponent(jScrollPane5);

        jSplitPane1.setRightComponent(jSplitPane2);

        TagList.addListSelectionListener(new javax.swing.event.ListSelectionListener() {
            public void valueChanged(javax.swing.event.ListSelectionEvent evt) {
                TagListValueChanged(evt);
            }
        });
        jScrollPane2.setViewportView(TagList);

        jSplitPane1.setLeftComponent(jScrollPane2);

        javax.swing.GroupLayout DataPanelLayout = new javax.swing.GroupLayout(DataPanel);
        DataPanel.setLayout(DataPanelLayout);
        DataPanelLayout.setHorizontalGroup(
            DataPanelLayout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(jSplitPane1)
        );
        DataPanelLayout.setVerticalGroup(
            DataPanelLayout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(jSplitPane1, javax.swing.GroupLayout.DEFAULT_SIZE, 411, Short.MAX_VALUE)
        );

        File_label.setText("Please select file from menu.");

        eventnumbersearchfield.setFormatterFactory(new javax.swing.text.DefaultFormatterFactory(new javax.swing.text.NumberFormatter(new java.text.DecimalFormat("#0"))));
        eventnumbersearchfield.setToolTipText("Event number");
        eventnumbersearchfield.addKeyListener(new java.awt.event.KeyAdapter() {
            public void keyPressed(java.awt.event.KeyEvent evt) {
                eventnumbersearchfieldKeyPressed(evt);
            }
        });

        eventlabel.setText("Event:");

        javax.swing.GroupLayout jPanel2Layout = new javax.swing.GroupLayout(jPanel2);
        jPanel2.setLayout(jPanel2Layout);
        jPanel2Layout.setHorizontalGroup(
            jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel2Layout.createSequentialGroup()
                .addComponent(eventlabel, javax.swing.GroupLayout.PREFERRED_SIZE, 54, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(eventnumbersearchfield, javax.swing.GroupLayout.PREFERRED_SIZE, 129, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addGap(18, 18, 18)
                .addComponent(File_label, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
        );
        jPanel2Layout.setVerticalGroup(
            jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel2Layout.createSequentialGroup()
                .addGroup(jPanel2Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.BASELINE)
                    .addComponent(eventnumbersearchfield, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(File_label, javax.swing.GroupLayout.PREFERRED_SIZE, 29, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(eventlabel))
                .addContainerGap(javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
        );

        PrevButton.setIcon(new javax.swing.ImageIcon("/home/chuwyjr/Downloads/icons8-long-arrow-left-32.png")); // NOI18N
        PrevButton.setText("Prev");
        PrevButton.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                PrevButtonActionPerformed(evt);
            }
        });

        NextButton.setIcon(new javax.swing.ImageIcon("/home/chuwyjr/Downloads/icons8-right-32.png")); // NOI18N
        NextButton.setText("Next");
        NextButton.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                NextButtonActionPerformed(evt);
            }
        });

        numbereventslabel.setText("Number of Events :");

        eventnumlabel.setText("#");

        entryidtext.setText("Entry ID");
        entryidtext.addKeyListener(new java.awt.event.KeyAdapter() {
            public void keyPressed(java.awt.event.KeyEvent evt) {
                entryidtextKeyPressed(evt);
            }
        });

        EnterButton.setText("Enter");
        EnterButton.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                EnterButtonActionPerformed(evt);
            }
        });

        javax.swing.GroupLayout jPanel4Layout = new javax.swing.GroupLayout(jPanel4);
        jPanel4.setLayout(jPanel4Layout);
        jPanel4Layout.setHorizontalGroup(
            jPanel4Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel4Layout.createSequentialGroup()
                .addComponent(numbereventslabel, javax.swing.GroupLayout.PREFERRED_SIZE, 140, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(eventnumlabel, javax.swing.GroupLayout.PREFERRED_SIZE, 123, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addGap(32, 32, 32)
                .addComponent(PrevButton, javax.swing.GroupLayout.PREFERRED_SIZE, 166, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(NextButton, javax.swing.GroupLayout.PREFERRED_SIZE, 170, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addGap(53, 53, 53)
                .addComponent(entryidtext, javax.swing.GroupLayout.PREFERRED_SIZE, 154, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addGap(33, 33, 33)
                .addComponent(EnterButton, javax.swing.GroupLayout.PREFERRED_SIZE, 69, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addContainerGap(javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
        );
        jPanel4Layout.setVerticalGroup(
            jPanel4Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel4Layout.createSequentialGroup()
                .addGroup(jPanel4Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.BASELINE)
                    .addComponent(entryidtext)
                    .addComponent(EnterButton))
                .addGap(13, 13, 13))
            .addGroup(jPanel4Layout.createSequentialGroup()
                .addGroup(jPanel4Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.BASELINE)
                    .addComponent(NextButton, javax.swing.GroupLayout.PREFERRED_SIZE, 57, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(PrevButton, javax.swing.GroupLayout.PREFERRED_SIZE, 56, javax.swing.GroupLayout.PREFERRED_SIZE))
                .addContainerGap(javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
            .addGroup(javax.swing.GroupLayout.Alignment.TRAILING, jPanel4Layout.createSequentialGroup()
                .addGap(6, 6, 6)
                .addGroup(jPanel4Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.BASELINE)
                    .addComponent(numbereventslabel, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                    .addComponent(eventnumlabel, javax.swing.GroupLayout.PREFERRED_SIZE, 40, javax.swing.GroupLayout.PREFERRED_SIZE))
                .addContainerGap())
        );

        entriesLabel.setText("Entries");

        tagsLabel.setText("Tags");

        entrydatalabel.setText("Entry Data");

        javax.swing.GroupLayout jPanel3Layout = new javax.swing.GroupLayout(jPanel3);
        jPanel3.setLayout(jPanel3Layout);
        jPanel3Layout.setHorizontalGroup(
            jPanel3Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(jPanel3Layout.createSequentialGroup()
                .addComponent(tagsLabel, javax.swing.GroupLayout.PREFERRED_SIZE, 268, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(entriesLabel, javax.swing.GroupLayout.PREFERRED_SIZE, 448, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED, 21, Short.MAX_VALUE)
                .addComponent(entrydatalabel, javax.swing.GroupLayout.PREFERRED_SIZE, 265, javax.swing.GroupLayout.PREFERRED_SIZE))
        );
        jPanel3Layout.setVerticalGroup(
            jPanel3Layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(entriesLabel, javax.swing.GroupLayout.PREFERRED_SIZE, 28, javax.swing.GroupLayout.PREFERRED_SIZE)
            .addComponent(entrydatalabel, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
            .addComponent(tagsLabel, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
        );

        jMenuBar1.setBorder(javax.swing.BorderFactory.createTitledBorder(""));

        File_Menu.setText("File");

        File_button.setSelected(true);
        File_button.setText("File");
        File_button.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                File_buttonActionPerformed(evt);
            }
        });
        File_Menu.add(File_button);

        jMenuBar1.add(File_Menu);

        Meta_menu.setText("Meta Data");

        HeaderRecord_button.setSelected(true);
        HeaderRecord_button.setText("Header Record");
        HeaderRecord_button.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                HeaderRecord_buttonActionPerformed(evt);
            }
        });
        Meta_menu.add(HeaderRecord_button);

        ParticalData_button.setSelected(true);
        ParticalData_button.setText("Particle Data");
        Meta_menu.add(ParticalData_button);

        Statistics_button.setSelected(true);
        Statistics_button.setText("Statistics");
        Meta_menu.add(Statistics_button);

        Logfile_button.setSelected(true);
        Logfile_button.setText("Logfile");
        Logfile_button.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                Logfile_buttonActionPerformed(evt);
            }
        });
        Meta_menu.add(Logfile_button);

        jMenuBar1.add(Meta_menu);

        DataLayout_menu.setText("Data Layout");

        DesLay_button.setSelected(true);
        DesLay_button.setText("Description Layout");
        DataLayout_menu.add(DesLay_button);

        HeadLay_button.setSelected(true);
        HeadLay_button.setText("Header layout");
        DataLayout_menu.add(HeadLay_button);

        EventLay_button.setSelected(true);
        EventLay_button.setText("Event Layout");
        DataLayout_menu.add(EventLay_button);

        StatLay_button.setSelected(true);
        StatLay_button.setText("Statistical Layout");
        DataLayout_menu.add(StatLay_button);

        jMenuBar1.add(DataLayout_menu);

        Help_menu.setText("Help");

        About_button.setSelected(true);
        About_button.setText("About");
        Help_menu.add(About_button);

        jMenuBar1.add(Help_menu);

        setJMenuBar(jMenuBar1);

        javax.swing.GroupLayout layout = new javax.swing.GroupLayout(getContentPane());
        getContentPane().setLayout(layout);
        layout.setHorizontalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(jPanel2, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
            .addComponent(jPanel4, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
            .addComponent(DataPanel, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
            .addComponent(jPanel3, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
        );
        layout.setVerticalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(layout.createSequentialGroup()
                .addComponent(jPanel2, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(jPanel4, javax.swing.GroupLayout.PREFERRED_SIZE, 58, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(jPanel3, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(DataPanel, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE))
        );

        pack();
    }                        

    private void File_buttonActionPerformed(java.awt.event.ActionEvent evt) {                                            
        JFileChooser chooser = new JFileChooser();
        chooser.showOpenDialog(null);
        f = chooser.getSelectedFile();
        filename = f.getAbsolutePath();

        File_label.setText(filename);

        try {
            reader = new Reader(filename);

            int nEvents = 0;
            for (Event event : reader) {

                nEvents++;
            }
            eventnumlabel.setText(String.valueOf(nEvents));

        } catch (Throwable e) {

        }
    }                                           

    private void eventnumbersearchfieldKeyPressed(java.awt.event.KeyEvent evt) {                                                  
        if (evt.getKeyCode() != java.awt.event.KeyEvent.VK_ENTER) {
            return;
        }
        long v = Long.parseLong(this.eventnumbersearchfield.getText());
        long p = Long.parseLong(eventnumlabel.getText());
        DefaultListModel model = new DefaultListModel();
        if (p - 1 == v) {
            long z = v - 1;
            eof(z);
        } else if (v >= 0 && v <= p) {
            try {
                reader.seekToStart();
                reader.skip(v);

                event = reader.next();
                for (String tag : event.getTags()) {
                    model.addElement(tag);
                }
            } catch (Throwable e) {
            }
        } else {
            JFrame frame = null;
            JOptionPane.showMessageDialog(frame, "Index out of Bounds");
        }

        this.TagList.setModel(model);
        this.TagList.revalidate();
        this.TagList.repaint();

    }                                                 

    private void EntryTableMouseClicked(java.awt.event.MouseEvent evt) {                                        
        JTable source = (JTable) evt.getSource();
        int row = source.rowAtPoint(evt.getPoint());
        Long id = Long.parseLong(source.getModel().getValueAt(row, 0) + "");
        if (id > 0) {
            try {
                Message entry = event.getEntry(id);
                this.EntryData.setText(entry.toString());
            } catch (Throwable e) {

            }
        }
    }                                       

    private void TagListValueChanged(javax.swing.event.ListSelectionEvent evt) {                                     
        if (evt.getValueIsAdjusting()) {
            return;
        }

        DefaultTableModel model = (DefaultTableModel) EntryTable.getModel();
        model.setRowCount(0);

        String tag = this.TagList.getSelectedValue();
        if (tag != null) {
            try {
                for (long entryID : event.getTaggedEntries(tag)) {
                    //model.addRow(new Object[]{entryID, event.getEntry(entryID).getClass()});

                    String es = new String(event.getEntry(entryID).getClass().toString());
                    String est = es.replace("$", ".");
                    String estr = est.replace("class", "");
                    String g = estr;
                    
                    List newlist = new ArrayList();
                    for (String word : Arrays.asList(g.split("\\."))) {

                        newlist.add(StringUtils.capitalize(word));
                    }

                    model.addRow(new Object[]{entryID, newlist});

                }
            } catch (Throwable e) {

            }
        }

        this.EntryTable.revalidate();
        this.EntryTable.repaint();
        this.EntryData.setText(null);
    }                                    

    private void Logfile_buttonActionPerformed(java.awt.event.ActionEvent evt) {                                               
        JFrame window = new JFrame();
        JTextArea text = new JTextArea();
        String h = "";
        String g = "";
        try {
            reader = new Reader(filename);
        } catch (IOException ex) {
            Logger.getLogger(proio_browser.class.getName()).log(Level.SEVERE, null, ex);
        }
        Event storeevent = null;
        Event tmpEvent;
        while (true) {
            tmpEvent = null;
            try {
                tmpEvent = reader.next(true);
            } catch (Throwable e) {
                ;
            }
            if (tmpEvent == null) {
                break;
            }
            storeevent = tmpEvent;
        }
        Map<String, ByteString> metadata = storeevent.getMetadata();
        for (Map.Entry<String, ByteString> entry : metadata.entrySet()) {
            if (entry.getValue().size() < 1000) {
                g = ("Metadata: " + entry.getKey() + ": " + entry.getValue().toStringUtf8() + "\n");
                h += g;
            }
        }

        text.setText(h);
        window.setLocationRelativeTo(null);
        window.setSize(720, 600);
        window.setTitle("Logfile");

        JScrollPane scroll = new JScrollPane();
        scroll.getViewport().setView(text);

        window.setContentPane(scroll);
        window.setVisible(true);

    }                                              

    private void HeaderRecord_buttonActionPerformed(java.awt.event.ActionEvent evt) {                                                    
        // TODO add your handling code here:
    }                                                   

    private void NextButtonActionPerformed(java.awt.event.ActionEvent evt) {                                           
        // TODO add your handling code here:
        String i = eventnumbersearchfield.getText();
        long t = Long.parseLong(i);

        long p = Long.parseLong(eventnumlabel.getText());
        long q = t + 1;
        if (p - 1 == q) {
            eof(q - 1);
        } else if (q >= p) {

            q = p;
            JFrame frame = null;
            JOptionPane.showMessageDialog(frame, "End of the file.");
        }
        String g = String.valueOf(q);

        eventnumbersearchfield.setText(g);
        this.jFormattedTextField1KeyPressed(true);

    }                                          

    private void PrevButtonActionPerformed(java.awt.event.ActionEvent evt) {                                           
        // TODO add your handling code here:
        String i = eventnumbersearchfield.getText();
        int t = Integer.parseInt(i);
        int q = t - 1;
        if (q < 0) {
            q = 0;
            JFrame frame = null;
            JOptionPane.showMessageDialog(frame, "You are at the beginning");
        }
        String g = new Integer(q).toString();

        eventnumbersearchfield.setText(g);
        jFormattedTextField1KeyPressed(true);
    }                                          

    private void EnterButtonActionPerformed(java.awt.event.ActionEvent evt) {                                            
       
        long id=Integer.parseInt(this.entryidtext.getText());
        JTable source =this.EntryTable;

        if (id > 0) {
            try {
                Message entry = event.getEntry(id);
                this.EntryData.setText(entry.toString());
            } catch (Throwable e) {

            }
        }
         for(int i = 0; i < source.getRowCount(); i++){//For each row
        for(int j = 0; j < source.getColumnCount(); j++){//For each column in that row
            if(source.getModel().getValueAt(i, j).equals(id)){//Search the model
                source.setRowSelectionInterval(i,i);
            }
        }//For loop inner
    }//For loop outer
        
    }                                           

    private void entryidtextKeyPressed(java.awt.event.KeyEvent evt) {                                       
         if (evt.getKeyCode() != java.awt.event.KeyEvent.VK_ENTER) {
            return;
        }
         EnterButton.doClick();
    }                                      

    /**
     * @param args the command line arguments
     */
    public static void main(String args[]) {
        /* Set the Nimbus look and feel */
        try {
            for (javax.swing.UIManager.LookAndFeelInfo info : javax.swing.UIManager.getInstalledLookAndFeels()) {
                if ("Nimbus".equals(info.getName())) {
                    javax.swing.UIManager.setLookAndFeel(info.getClassName());
                    break;
                }
            }
        } catch (ClassNotFoundException ex) {
            java.util.logging.Logger.getLogger(proio_browser.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (InstantiationException ex) {
            java.util.logging.Logger.getLogger(proio_browser.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (IllegalAccessException ex) {
            java.util.logging.Logger.getLogger(proio_browser.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (javax.swing.UnsupportedLookAndFeelException ex) {
            java.util.logging.Logger.getLogger(proio_browser.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        }

        java.awt.EventQueue.invokeLater(new Runnable() {
            @Override
            public void run() {
                new proio_browser().setVisible(true);
            }
        });
    }

    // Variables declaration - do not modify                     
    private javax.swing.JRadioButtonMenuItem About_button;
    private javax.swing.JMenu DataLayout_menu;
    private javax.swing.JPanel DataPanel;
    private javax.swing.JRadioButtonMenuItem DesLay_button;
    private javax.swing.JButton EnterButton;
    private javax.swing.JTextArea EntryData;
    private javax.swing.JTable EntryTable;
    private javax.swing.JRadioButtonMenuItem EventLay_button;
    private javax.swing.JMenu File_Menu;
    private javax.swing.JRadioButtonMenuItem File_button;
    private javax.swing.JLabel File_label;
    private javax.swing.JRadioButtonMenuItem HeadLay_button;
    private javax.swing.JRadioButtonMenuItem HeaderRecord_button;
    private javax.swing.JMenu Help_menu;
    private javax.swing.JRadioButtonMenuItem Logfile_button;
    private javax.swing.JMenu Meta_menu;
    private javax.swing.JButton NextButton;
    private javax.swing.JRadioButtonMenuItem ParticalData_button;
    private javax.swing.JButton PrevButton;
    private javax.swing.JRadioButtonMenuItem StatLay_button;
    private javax.swing.JRadioButtonMenuItem Statistics_button;
    private javax.swing.JList<String> TagList;
    private java.awt.Label entriesLabel;
    private java.awt.Label entrydatalabel;
    private javax.swing.JTextField entryidtext;
    private javax.swing.JLabel eventlabel;
    private javax.swing.JFormattedTextField eventnumbersearchfield;
    private javax.swing.JLabel eventnumlabel;
    private javax.swing.JMenuBar jMenuBar1;
    private javax.swing.JPanel jPanel2;
    private javax.swing.JPanel jPanel3;
    private javax.swing.JPanel jPanel4;
    private javax.swing.JScrollPane jScrollPane1;
    private javax.swing.JScrollPane jScrollPane2;
    private javax.swing.JScrollPane jScrollPane5;
    private javax.swing.JSplitPane jSplitPane1;
    private javax.swing.JSplitPane jSplitPane2;
    private javax.swing.JLabel numbereventslabel;
    private java.awt.Label tagsLabel;
    // End of variables declaration                   
    public File f;
    public String filename;
    public Reader reader;
    private Event event;

    private void jFormattedTextField1KeyPressed(boolean b) {
        long v = Long.parseLong(this.eventnumbersearchfield.getText());

        DefaultListModel model = new DefaultListModel();
        if (v >= 0) {
            try {
                reader.seekToStart();
                reader.skip(v);

                event = reader.next();
                for (String tag : event.getTags()) {
                    model.addElement(tag);
                }
            } catch (Throwable e) {
                ;
            }
        }

        this.TagList.setModel(model);
        this.TagList.revalidate();
        this.TagList.repaint();
    }

    private void eof(long par) {
        String g = "";
        String h = "";

        try {
            reader = new Reader(filename);
        } catch (IOException ex) {
            Logger.getLogger(proio_browser.class.getName()).log(Level.SEVERE, null, ex);
        }
        Event storeevent = null;
        Event tmpEvent;
        while (true) {
            tmpEvent = null;
            try {
                tmpEvent = reader.next(true);
            } catch (Throwable e) {
                ;
            }
            if (tmpEvent == null) {
                break;
            }
            storeevent = tmpEvent;
        }
        Map<String, ByteString> metadata = storeevent.getMetadata();
        for (Map.Entry<String, ByteString> entry : metadata.entrySet()) {
            if (entry.getValue().size() < 1000) {
                g = ("Metadata: " + entry.getKey() + ": " + entry.getValue().toStringUtf8() + "\n");
                h += g;
            }
        }

        EntryData.setText(h);
        EntryData.revalidate();
        EntryData.repaint();
    }
}
