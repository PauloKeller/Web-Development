select * from cd.facilities; 

select name, membercost from cd.facilities;

select * from cd.facilities where membercost > 0;

select facid, name, membercost, monthlymaintenance from cd.facilities where membercost > 0 and (membercost < monthlymaintenance/50.0);

select * from cd.facilities where name like '%Tennis%';

select * from cd.facilities where facid in (1,5);

select memid, surname, firstname, joindate from cd.members where joindate >= '2012-09-01';

select distinct surname from cd.members order by surname limit 10;

select max(joindate) as latest from cd.members;

select count(*) from cd.facilities where guestcost >= 10;

select facid, sum(slots) as "Total Slots" from cd.bookings where starttime >= '2012-09-01' and starttime < '2012-10-01' group by facid order by sum(slots);

select facid, sum(slots) as "Total Slots" from cd.bookings group by facid having sum(slots) > 1000 order by facid;

select bks.starttime as start, facs.name as name from cd.facilities facs inner join cd.bookings bks on facs.facid = bks.facid where facs.facid in (0,1) and bks.starttime >= '2012-09-21' and bks.starttime < '2012-09-22' order by bks.starttime;

select bks.starttime from cd.bookings bks inner join cd.members mems on mems.memid = bks.memid where mems.firstname='David' and mems.surname='Farrell';