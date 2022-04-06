import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, DeleteButton, EditButton, Create, DateInput} from 'react-admin';
  
  export const CommandsList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="firstname" label="Prénom"/>
        <TextField source="lastname" label="Nom"/>
        <TextField source="id_user" label="Id utilisateur"/>
        <TextField source="price" label="Prix"/>
        <TextField source="quantity" label="Quantité"/>
        <TextField source="date" label="Date"/>
        <TextField source="status" label="Statut"/>
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const CommandsEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="status" label="Statut"/>
      </SimpleForm>
    </Edit>
  );

  export const CommandsCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="id_user" label="Id utilisateur"/>
        <DateInput source="date" label="Date"/>
        <TextInput source="status" label="Statut"/>
      </SimpleForm>
    </Create>
  );
  