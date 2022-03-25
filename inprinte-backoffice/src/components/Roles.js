import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, Create,} from 'react-admin';
  
  export const RolesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="name" />
      </Datagrid>
    </List>
  );
  
  export const RolesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="id" disabled />
        <TextInput source="name" />
      </SimpleForm>
    </Edit>
  );

  export const RolesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="name" />
      </SimpleForm>
    </Create>
  );