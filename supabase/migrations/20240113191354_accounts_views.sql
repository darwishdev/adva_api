CREATE OR REPLACE VIEW accounts_schema.users_view AS
WITH userpermissions AS (
  SELECT 
    ur.user_id, 
    rp.permission_id 
  FROM 
    accounts_schema.user_roles ur 
    JOIN accounts_schema.role_permissions rp ON rp.role_id = ur.role_id 
  ORDER BY 
    ur.user_id
), 
permissionsgroups AS (
  SELECT 
    p.permission_group 
  FROM 
    accounts_schema.permissions p 
  GROUP BY 
    p.permission_group
) ,userRecord as (
SELECT 
  u.user_id, 
  user_name, 
  user_image, 
  user_email, 
  user_phone,  
  user_password, 
  u.created_at, 
  u.updated_at, 
  Array_agg(up.permission_id) permissions, 
  u.deleted_at, 
  (
    SELECT 
      Jsonb_agg(nested_groups) groups 
    FROM 
      (
        SELECT 
          g.permission_group, 
          (
            SELECT 
              Jsonb_agg(nested_permissions) 
            FROM 
              (
                SELECT 
                  np.permission_function, 
                  (up.permission_id IS NOT NULL) authorized 
                FROM 
                  accounts_schema.permissions np 
                  LEFT JOIN userpermissions up ON np.permission_id = up.permission_id 
                  AND up.user_id = u.user_id 
                WHERE 
                  np.permission_group = g.permission_group
              ) nested_permissions
          ) permissions 
        FROM 
          permissionsgroups g
      ) nested_groups
  ) AS permission_groups 
FROM 
  accounts_schema.users u 
  JOIN userpermissions up ON u.user_id = up.user_id 
  GROUP BY 
  u.user_id, 
  user_name, 
    user_image, 
  user_email, 
   user_phone, 
  user_password, 
  u.created_at, 
  u.updated_at, 
  u.deleted_at

) 
select
ur.user_id,
ur.user_name,
 
 
 
 ur.user_image,
ur.user_email,
ur.user_phone,
 ur.user_password,
ur.created_at,
ur.updated_at,
ur.permissions,
ur.deleted_at ,
ur.permission_groups,
(
        SELECT
            Jsonb_Agg(nested_sidebar) side_bar
        FROM (
           SELECT
        navigation_bar_id,
        menu_key "key",
        label,
        icon,
        "to",
         (SELECT
                Jsonb_agg(nested_items)
            FROM (
               SELECT
                    child.parent_id,
                    child.menu_key "key",
                    child.label,
                    child.icon,
                    child."to"
                FROM
                    accounts_schema.navigation_bars child
                WHERE
                
                    child.parent_id = parent.navigation_bar_id and
                    child.permission_id = ANY (
                        ur.permissions)
                    ORDER BY
                        child.menu_key
            )nested_items)items

    FROM
        accounts_schema.navigation_bars parent
    WHERE
        parent_id IS NULL
        AND ("to" IS NULL
            OR parent.permission_id = ANY (ur.permissions))
           ) nested_sidebar) side_bar
from userRecord ur ;