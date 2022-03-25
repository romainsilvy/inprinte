import {List, Datagrid, TextField, Edit, SimpleForm, TextInput} from 'react-admin';
  
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
      </Datagrid>
    </List>
  );
  
  export const CommandLinesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="id" disabled />
        <TextField source="status" />
      </SimpleForm>
    </Edit>
  );