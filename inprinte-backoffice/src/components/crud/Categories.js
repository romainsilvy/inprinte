import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, Create, EditButton, DeleteButton} from 'react-admin';
  
  export const CategoriesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="name" />
        <TextField source="is_alive" />
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
  
  export const CategoriesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="name" />
        <TextInput source="is_alive" />
      </SimpleForm>
    </Edit>
  );

  export const CategoriesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="name" />
        <TextInput source="is_alive" />
      </SimpleForm>
    </Create>
  );