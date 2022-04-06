import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, DeleteButton, EditButton, Create, DateInput} from 'react-admin';
  
  export const CommandsList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="firstname" />
        <TextField source="lastname" />
        <TextField source="id_user" />
        <TextField source="price" />
        <TextField source="quantity" />
        <TextField source="date" />
        <TextField source="status" />
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const CommandsEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="status" />
      </SimpleForm>
    </Edit>
  );

  export const CommandsCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="id_user" />
        <DateInput source="date" />
        <TextInput source="status" />
      </SimpleForm>
    </Create>
  );