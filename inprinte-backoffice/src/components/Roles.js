import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, Create, EditButton, DeleteButton} from 'react-admin';
  
  export const RolesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="role" />
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const RolesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="role" />
      </SimpleForm>
    </Edit>
  );

  export const RolesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="role" />
      </SimpleForm>
    </Create>
  );