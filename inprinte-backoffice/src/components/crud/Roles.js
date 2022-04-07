import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, Create, EditButton, DeleteButton} from 'react-admin';
  
  export const RolesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" label="Id"/>
        <TextField source="role" label="Role"/>
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const RolesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="role" label="Role"/>
      </SimpleForm>
    </Edit>
  );

  export const RolesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="role" label="Role"/>
      </SimpleForm>
    </Create>
  );