use seata_order;
delete from undo_log;
delete from so_a;
delete from branch_transaction;

use seata_product;
delete from undo_log;
delete from so_b;
delete from branch_transaction;

use seata;
delete from branch_table;
delete from global_table;
delete from lock_table;
