import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, EditButton, DeleteButton, Create} from 'react-admin';
  
  export const CommandLinesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" label="Id"/>
        <TextField source="id_user" label="Id utilisateur"/>
        <TextField source="firstname" label="Prénom"/>
        <TextField source="lastname" label="Nom"/>
        <TextField source="id_command" label="Id commande"/>
        <TextField source="id_product" label="Id produit"/>
        <TextField source="price" label="Prix"/>
        <TextField source="quantity" label="Quantité"/>
        <TextField source="date" label="Date"/>
        <TextField source="status" label="Statut"/>
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const CommandLinesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="status" label="Statut"/>
      </SimpleForm>
    </Edit>
  );

  export const CommandLinesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="id_product" label="Id produit"/>
        <TextInput source="id_command" label="Id commande"/>
        <TextInput source="quantity" label="Quantité"/>
        <TextInput source="status" label="Statut"/>
      </SimpleForm>
    </Create>
  );