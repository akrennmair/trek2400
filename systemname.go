package main

func systemname(q *quadrant) string {
	i := q.qsystemname
	if i&Q_DISTRESSED != 0 {
		i = eventList[i&Q_SYSTEM].systemname
	}

	i &= Q_SYSTEM
	if i == 0 {
		return ""
	}
	return systemnameList[i]
}
