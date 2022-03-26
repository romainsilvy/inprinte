import {List, Datagrid, TextField, Edit, SimpleForm, TextInput} from 'react-admin';
  
  export const CommandsList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" label="id_command"/>
        <TextField source="firstname" />
        <TextField source="lastname" />
        <TextField source="id_user" />
        <TextField source="price" />
        <TextField source="quantity" />
        <TextField source="date" />
        <TextField source="status" />
      </Datagrid>
    </List>
  );
  
  export const CommandsEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="id" disabled />
        <TextField source="status" />
      </SimpleForm>
    </Edit>
  );