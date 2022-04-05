import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, EditButton, DeleteButton, Create} from 'react-admin';
  
  export const CommandLinesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="id_user" />
        <TextField source="firstname" />
        <TextField source="lastname" />
        <TextField source="id_command" />
        <TextField source="id_product" />
        <TextField source="price" />
        <TextField source="quantity" />
        <TextField source="date" />
        <TextField source="status" />
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const CommandLinesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="status" />
      </SimpleForm>
    </Edit>
  );

  export const CommandLinesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="id_product" />
        <TextInput source="id_command" />
        <TextInput source="quantity" />
        <TextInput source="status" />
      </SimpleForm>
    </Create>
  );